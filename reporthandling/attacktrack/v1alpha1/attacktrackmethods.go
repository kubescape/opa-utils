package v1alpha1

import (
	"github.com/kubescape/k8s-interface/workloadinterface"
)

func (at *AttackTrack) GetApiVersion() string {
	return at.ApiVersion
}

func (at *AttackTrack) GetKind() string {
	return at.Kind
}

func (at *AttackTrack) GetName() string {
	if v, ok := workloadinterface.InspectMap(at.Metadata, "name"); ok {
		return v.(string)
	}
	return ""
}

func (at *AttackTrack) GetDescription() string {
	return at.Spec.Description
}

func (at *AttackTrack) GetVersion() string {
	return at.Spec.Version
}

func (at *AttackTrack) GetData() IAttackTrackStep {
	return &at.Spec.Data
}

// IsValid returns true if an attack track is valid
func (at *AttackTrack) IsValid() bool {
	// A valid AttackTrack must be a tree (Directed Acyclic Graph where a child node can only have one parent)

	// We validate the AttackTrack by performing a DFS on the AttackTrack definition and checking if we visit a node twice, if we do, then the graph is cyclic (invalid).
	// In case a child node has multiple parents, we will visit it multiple times (because we run DFS on all children of a node).
	// According to the definition of the AttackTrack, we can only have a single root node, so we can start the DFS from that node.
	// All nodes are reachable from the root node, per the definition of the AttackTrack, so our graph is connected.

	visited := make(map[string]bool)
	return directedDfs(at.GetData(), visited)
}

// ==================== Iterator ====================
func (at *AttackTrack) Iterator() IAttackTrackIterator {
	s := &attackTrackStepStack{}
	s.Push(at.GetData())

	return &AttackTrackIterator{
		stack: s,
	}
}

func (iter *AttackTrackIterator) HasNext() bool {
	return !iter.stack.IsEmpty()
}

func (iter *AttackTrackIterator) Next() IAttackTrackStep {
	step := iter.stack.Pop()
	for i := 0; i < step.Length(); i++ {
		iter.stack.Push(step.SubStepAt(i))
	}

	return step
}

func NewAttackTrackAllPathsHandler(attackTrack IAttackTrack, lookup IAttackTrackControlsLookup) *AttackTrackAllPathsHandler {
	inDegreeZero := make(map[string]bool)
	outDegreeZero := make(map[string]bool)

	// initialize all nodes with in-degree 0 and out-degree 0
	iter := attackTrack.Iterator()
	for iter.HasNext() {
		step := iter.Next()
		inDegreeZero[step.GetName()] = true
		outDegreeZero[step.GetName()] = true

		attackTrackName := attackTrack.GetName()
		// load failed controls for each step
		failedControls := lookup.GetAssociatedControls(attackTrackName, step.GetName())
		step.SetControls(failedControls)
	}

	adjacencyMatrix := make(map[string][]IAttackTrackStep)
	visited := make(map[string]bool)

	return &AttackTrackAllPathsHandler{
		attackTrack:     attackTrack,
		inDegreeZero:    inDegreeZero,
		outDegreeZero:   outDegreeZero,
		adjacencyMatrix: adjacencyMatrix,
		visited:         visited,
	}
}

// calculateAdjacencyMatrixAndVerticesDegree calculates the adjacency matrix and the in and out degree of each vertex in the graph
func (h *AttackTrackAllPathsHandler) calculateAdjacencyMatrixAndVerticesDegree(u IAttackTrackStep, v IAttackTrackStep) {
	// u -> v (u is the parent of v)
	if u.IsPartOfAttackTrackPath() && v.IsPartOfAttackTrackPath() {
		h.adjacencyMatrix[u.GetName()] = append(h.adjacencyMatrix[u.GetName()], v)
		h.inDegreeZero[v.GetName()] = false
		h.outDegreeZero[u.GetName()] = false
	}

	for i := 0; i < v.Length(); i++ {
		h.calculateAdjacencyMatrixAndVerticesDegree(v, v.SubStepAt(i))
	}
}

func (h *AttackTrackAllPathsHandler) allPathsDfs(allPaths [][]IAttackTrackStep, step IAttackTrackStep, currentPathIndex *int) [][]IAttackTrackStep {
	allPaths[*currentPathIndex] = append(allPaths[*currentPathIndex], step)
	h.visited[step.GetName()] = true

	// if the current step is a leaf, then we have found a path and path is complete
	if h.outDegreeZero[step.GetName()] && h.inDegreeZero[allPaths[*currentPathIndex][0].GetName()] {
		newPath := make([]IAttackTrackStep, len(allPaths[*currentPathIndex]))
		copy(newPath, allPaths[*currentPathIndex])
		allPaths = append(allPaths, newPath)
		*currentPathIndex++
	}

	for i := range h.adjacencyMatrix[step.GetName()] {
		v := h.adjacencyMatrix[step.GetName()][i]
		if !v.IsPartOfAttackTrackPath() {
			continue
		}

		if !h.visited[v.GetName()] {
			allPaths = h.allPathsDfs(allPaths, v, currentPathIndex)
		}
	}

	// Backtrack
	allPaths[*currentPathIndex] = allPaths[*currentPathIndex][:len(allPaths[*currentPathIndex])-1]
	h.visited[step.GetName()] = false
	return allPaths
}

func (handler *AttackTrackAllPathsHandler) CalculateAllPaths() [][]IAttackTrackStep {
	// calculate the adjacency matrix of the attack track and the in-degree and out-degree of each node
	root := handler.attackTrack.GetData()

	for i := 0; i < root.Length(); i++ {
		handler.calculateAdjacencyMatrixAndVerticesDegree(root, root.SubStepAt(i))
	}

	allPaths := make([][]IAttackTrackStep, 0)
	var currentPathIndex *int = new(int)
	*currentPathIndex = -1
	iter := handler.attackTrack.Iterator()
	for iter.HasNext() {
		step := iter.Next()
		if !step.IsPartOfAttackTrackPath() || !handler.inDegreeZero[step.GetName()] {
			continue
		}

		(*currentPathIndex)++
		allPaths = append(allPaths, []IAttackTrackStep{})
		allPaths = handler.allPathsDfs(allPaths, step, currentPathIndex)

		// If last DFS call added a new empty path remove it
		if *currentPathIndex >= 0 && len(allPaths[*currentPathIndex]) == 0 {
			allPaths = allPaths[:len(allPaths)-1]
			(*currentPathIndex)--
		}
	}

	return allPaths
}

func NewAttackTrackControlsLookup(attackTracks []IAttackTrack, failedControlIds []string, allControls map[string]IAttackTrackControl) AttackTrackControlsLookup {
	lookup := make(AttackTrackControlsLookup)

	for _, attackTrack := range attackTracks {
		attackTrackName := attackTrack.GetName()

		lookup[attackTrackName] = make(map[string][]IAttackTrackControl)
		for _, controlId := range failedControlIds {
			control := allControls[controlId]
			for _, category := range control.GetAttackTrackCategories(attackTrackName) {
				lookup[attackTrackName][category] = append(lookup[attackTrackName][category], control)
			}
		}
	}

	return lookup
}

func (at *AttackTrackControlsLookup) GetAssociatedControls(attackTrack, category string) []IAttackTrackControl {
	if v, ok := (*at)[attackTrack][category]; ok {
		return v
	}
	return make([]IAttackTrackControl, 0)
}

func (at *AttackTrackControlsLookup) HasAssociatedControls(attackTrack string) bool {
	if controls, ok := (*at)[attackTrack]; ok && len(controls) > 0 {
		return true
	}
	return false
}

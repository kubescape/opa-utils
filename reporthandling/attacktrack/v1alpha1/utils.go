package v1alpha1

type attackTrackStepStack []IAttackTrackStep

func directedDfs(node IAttackTrackStep, visited map[string]bool) bool {
	if _, ok := visited[node.GetName()]; ok {
		return false
	}

	visited[node.GetName()] = true

	for i := 0; i < node.Length(); i++ {
		if !directedDfs(node.SubStepAt(i), visited) {
			return false
		}
	}
	return true
}

func (s *attackTrackStepStack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *attackTrackStepStack) Push(node IAttackTrackStep) {
	*s = append(*s, node)
}

func (s *attackTrackStepStack) Pop() IAttackTrackStep {
	if s.IsEmpty() {
		return nil
	}

	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element
}

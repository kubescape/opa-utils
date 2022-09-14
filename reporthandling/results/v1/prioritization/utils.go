package prioritization

import "github.com/kubescape/opa-utils/reporthandling/attacktrack/v1alpha1"

// ControlsVectorFromAttackTrackPaths creates a list of controls vectors from a list of attack track paths
func ControlsVectorFromAttackTrackPaths(attackTrack v1alpha1.IAttackTrack, paths [][]v1alpha1.IAttackTrackStep) []*ControlsVector {
	controlsVectors := make([]*ControlsVector, 0)
	for _, path := range paths {
		controlsVectors = controlsVectorFromSinglePath(attackTrack, path, controlsVectors)
	}
	return controlsVectors
}

// ControlsVectorFromSinglePath creates a list of controls vector from a single attack track path
//
// For example, if the attack track path is:
//
//		(A) 1,2 -> (B) 3,4 -> (C) 5
//	 Where letters are steps and numbers are controls in that step
//
//	 Then the returned list will be a list of controls vectors:
//
//		[0] (1,A) -> (3,B) -> (5,C)
//		[1] (1,A) -> (4,B) -> (5,C)
//		[2] (2,A) -> (3,B) -> (5,C)
//		[3] (2,A) -> (4,B) -> (5,C)
func controlsVectorFromSinglePath(attackTrack v1alpha1.IAttackTrack, path []v1alpha1.IAttackTrackStep, result []*ControlsVector) []*ControlsVector {
	// number of controls arrays
	n := len(path)

	// to keep track of next element
	// in each of the n controls arrays
	indices := make([]int, n)
	for i := range indices {
		indices[i] = 0
	}

	for {
		// create a new controls vector combination
		controlsVector := NewControlsVector(attackTrack.GetName())

		for i := 0; i < n; i++ {
			controls := path[i].GetControls()
			controlsVector.AddControl(
				&PriorityVectorControl{
					ControlID: controls[indices[i]].GetControlId(),
					Category:  path[i].GetName(),
					Tags:      controls[indices[i]].GetControlTypeTags(),
				})
		}

		if controlsVector.IsValid() {
			result = append(result, controlsVector)
		}

		// find the rightmost array that has more elements
		// left after the current element in that array
		next := n - 1
		for next >= 0 && (indices[next]+1 >= len(path[next].GetControls())) {
			next--
		}

		// No such array found so no more combinations left
		if next < 0 {
			return result
		}

		// If found, move to next element in that array
		indices[next]++

		// for all arrays to the right of this array
		// current index again points to first element
		for i := next + 1; i < n; i++ {
			indices[i] = 0
		}
	}
}

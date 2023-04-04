package main

func Contains(elems []int, v int) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func RemoveAtIndexKItem(slice []KItem, indices []int) []KItem {
	// Create a new slice to hold the filtered elements
	filtered := make([]KItem, 0)
	// Create a map to hold the indices to remove
	indicesToRemove := make(map[int]bool)
	for _, index := range indices {
		indicesToRemove[index] = true
	}
	// Iterate through the original slice, appending the elements that should not be removed
	for i, element := range slice {
		if !indicesToRemove[i] {
			filtered = append(filtered, element)
		}
	}
	return filtered
}

func RemoveAtIndexPackage(slice []Package, indices []int) []Package {
	// Create a new slice to hold the filtered elements
	filtered := make([]Package, 0)
	// Create a map to hold the indices to remove
	indicesToRemove := make(map[int]bool)
	for _, index := range indices {
		indicesToRemove[index] = true
	}
	// Iterate through the original slice, appending the elements that should not be removed
	for i, element := range slice {
		if !indicesToRemove[i] {
			filtered = append(filtered, element)
		}
	}
	return filtered
}

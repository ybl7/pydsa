package main

// Return a slice with all indexes where all indexes right of it are less
// Must be O(n) since we aren't searching, or concerned with endpoints, we need to inspect all elements
// Makes sense to iterate the array backwards and keep track of the maximum value from the back, only update it
// when we come across a larger value. The question asks for a sorted output so we can reverse in O(n) time at the end.
func BuildingOceanView(arr []int) []int {
	// Rightmost building always has an ocean view
	out := []int{
		arr[len(arr)-1],
	}

	for i, j := len(arr)-2, len(arr)-1; i > -1; i-- {
		if arr[i] > arr[j] {
			out = append(out, i)
			j = i
		}
	}

	for k, l := 0, len(out)-1; k < l; k, l = k+1, l-1 {
		out[k], out[l] = out[l], out[k]
	}

	return out
}

// [1,4,2,3]
// j = 3, i = 2
// MaxFromBack = 3
// arr[i] < arr[j] so no view, don't update j

// [1,4,2,3]
// j = 3, i = 1
// MaxFromBack = 3
// arr[i] > arr[j] so view, add i to out, update j to i

package main

// [7,1,5,3,6,4]
// [7,6,4,3,1]
func BuySellStock(arr []int) int {
	max := 0

	for i, j := 0, 1; j < len(arr); j++ {
		curr := arr[j] - arr[i]
		if curr > max {
			max = curr
		}

		// Update i if we find a better lower bound
		if arr[j] < arr[i] {
			i = j
		}
	}
	return max
}

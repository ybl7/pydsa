package main

func MaxSumSubarr(arr []int) int {
	endAt := arr[0]
	maxV := arr[0]

	for i := 1; i < len(arr); i++ {
		endAt = min(arr[i], endAt+arr[i])
		maxV = max(maxV, endAt)
	}

	return maxV
}

package main

func PeakElement(arr []int) int {
	l, r := 0, len(arr)-1

	for l < r {
		m := (r + l) / 2
		if arr[m] < arr[m+1] {
			// discard left side, since arr[m+1] > arr[m], it's our candidate for a peak
			l = m + 1
		} else if arr[m] > arr[m+1] {
			// discard right side, since arr[m] > arr[m+1], it's our candidate for a peak
			r = m
		}
	}
	// we return l because the exist condition is when l == r since l only increases and r only decreases
	// since we proved that the interval that we keep MUST have a peak in it, if l == r then this 0 width interval must be a peak
	return l
}

// Example [1, 2, 3, 1]

// l = 0, r = 3, m = 1
// arr[m] = 2 which is < 3, so 3 is our current peak

package main

import (
	"sort"
)

func main()

func MergeIntervals(intervals [][]int) [][]int {
	// After sorting we guarantee incoming intervals always have greater start than current intervals
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	out := [][]int{}
	out = append(out, intervals[0])

	for i := 1; i < len(intervals); i++ {
		currEnd := out[len(out)-1][1]
		incStart := intervals[i][0]
		incEnd := intervals[i][1]

		// Incoming interval does not overlap with current interval
		if incStart > currEnd {
			out = append(out, []int{incStart, incEnd})
		} else {
			// Incoming interval does overlap, take the max of the ending times
			if incEnd > currEnd {
				out[len(out)-1][1] = incEnd
			}
		}
	}

	return out
}

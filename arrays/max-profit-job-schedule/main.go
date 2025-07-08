package main

import (
	"sort"
)

func main() {
	MaxProfitJobSched([]int{1, 2, 4, 6, 5, 7}, []int{3, 5, 6, 7, 8, 9}, []int{5, 6, 5, 4, 11, 2})
	MaxProfitJobSched([]int{1, 2, 3, 4, 6}, []int{3, 5, 10, 6, 9}, []int{20, 20, 100, 70, 60})
	MaxProfitJobSched([]int{1, 1, 1}, []int{2, 3, 4}, []int{5, 6, 4})
}

func MaxProfitJobSched(starts, ends, profits []int) int {
	jobs := []Job{}
	for idx, _ := range starts {
		jobs = append(jobs, Job{
			Start:  starts[idx],
			End:    ends[idx],
			Profit: profits[idx],
		})
	}

	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].End < jobs[j].End
	})

	profAtEnd := make([]int, len(jobs))
	for i, j := range jobs {
		profAtEnd[i] = j.Profit
	}

	for j := 1; j < len(profAtEnd); j++ {
		for i := 0; i < j; i++ {
			if jobs[i].End <= jobs[j].Start {
				// Don't overlap
				profAtEnd[j] = max(profAtEnd[j], jobs[j].Profit+profAtEnd[i])
			}
		}
	}

	max := 0
	for _, e := range profAtEnd {
		if max < e {
			max = e
		}
	}

	return max
}

type Job struct {
	Start  int
	End    int
	Profit int
}

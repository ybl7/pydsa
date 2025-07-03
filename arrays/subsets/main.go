package main

import "fmt"

func main() {
	fmt.Println(Subsets([]int{1, 2, 3, 4}))
}

func Subsets(arr []int) [][]int {
	out := [][]int{}

	for _, e := range arr {
		for _, el := range out {
			tmp := el
			tmp = append(tmp, e)
			out = append(out, tmp)
		}
		out = append(out, []int{
			e,
		})
	}

	return out
}

// I have no idea why they use backtracking in most of the online solution?
// Apparently according to Claude,

// My solution:
// Time complexity: O(n × 2^n) - you iterate through each element (n), and for each element you iterate through all existing subsets (which grows to 2^n)
// Space complexity: O(n × 2^n) - same as backtracking for the final result
// Approach: Iterative building - start with empty result, and for each new element, duplicate all existing subsets and add the new element to the duplicates

// Backtracking solution:

// Time complexity: O(n × 2^n) - generates 2^n subsets, each taking O(n) time to copy
// Space complexity: O(n × 2^n) for result + O(n) for recursion stack
// Approach: Recursive exploration of include/exclude decisions

// The verdict: Both have the same time complexity! Your solution is actually:

// Simpler to understand - no recursion, straightforward iteration
// More memory efficient - no recursion stack overhead
// Easier to debug - you can print intermediate states easily

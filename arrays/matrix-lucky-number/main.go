package main

import (
	"fmt"
	"reflect"
)

// Test cases LLM generated
func main() {
	tests := []struct {
		name     string
		matrix   [][]int
		expected []int
	}{
		{
			name: "Single lucky number example 1",
			matrix: [][]int{
				{3, 7, 8},
				{9, 11, 13},
				{15, 16, 17},
			},
			expected: []int{15},
		},
		{
			name: "Single lucky number example 2",
			matrix: [][]int{
				{1, 10, 4, 2},
				{9, 3, 8, 7},
				{15, 16, 17, 12},
			},
			expected: []int{12},
		},
		{
			name: "Multiple lucky numbers",
			matrix: [][]int{
				{7, 8},
				{1, 2},
			},
			expected: []int{7},
		},
		{
			name: "Lucky number 7",
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: []int{7},
		},
		{
			name: "Single element matrix",
			matrix: [][]int{
				{42},
			},
			expected: []int{42},
		},
		{
			name: "Single row matrix",
			matrix: [][]int{
				{1, 2, 3, 4, 5},
			},
			expected: []int{1},
		},
		{
			name: "Single column matrix",
			matrix: [][]int{
				{1},
				{2},
				{3},
				{4},
				{5},
			},
			expected: []int{5},
		},
		{
			name: "2x2 matrix with lucky number",
			matrix: [][]int{
				{3, 1},
				{4, 2},
			},
			expected: []int{2},
		},
		{
			name: "Mixed positive and negative - lucky number 6",
			matrix: [][]int{
				{-1, 5, 3},
				{0, 2, 4},
				{-3, 1, 6},
			},
			expected: []int{0},
		},
		{
			name: "All same values in row - lucky number 5",
			matrix: [][]int{
				{5, 5, 5},
				{1, 2, 3},
				{4, 6, 7},
			},
			expected: []int{5},
		},
		{
			name: "All same values in column",
			matrix: [][]int{
				{1, 5, 7},
				{2, 5, 8},
				{3, 5, 9},
			},
			expected: []int{3},
		},
		{
			name: "Diagonal pattern",
			matrix: [][]int{
				{9, 2, 3},
				{4, 5, 6},
				{1, 8, 7},
			},
			expected: []int{},
		},
		{
			name: "No lucky numbers - each min not max",
			matrix: [][]int{
				{3, 1, 2},
				{6, 4, 5},
				{9, 7, 8},
			},
			expected: []int{7},
		},
	}

	passedTests := 0
	totalTests := len(tests)

	fmt.Println("Running Lucky Numbers in Matrix Tests")
	fmt.Println("=====================================")

	for _, test := range tests {
		result := MatLuckyNum(test.matrix)
		passed := reflect.DeepEqual(result, test.expected)

		if passed {
			fmt.Printf("✓ PASS: %s\n", test.name)
			passedTests++
		} else {
			fmt.Printf("✗ FAIL: %s\n", test.name)
			fmt.Printf("  Input:    %v\n", test.matrix)
			fmt.Printf("  Expected: %v\n", test.expected)
			fmt.Printf("  Got:      %v\n", result)
		}
	}

	fmt.Println("=====================================")
	fmt.Printf("Tests passed: %d/%d\n", passedTests, totalTests)

	if passedTests == totalTests {
		fmt.Println("🎉 All tests passed!")
	} else {
		fmt.Printf("❌ %d tests failed\n", totalTests-passedTests)
	}
}

func MatLuckyNum(m [][]int) []int {
	out := []int{}
	nRows := len(m)

	for _, row := range m {
		// Minimum value of current row, and it's column
		colIdx, rMinVal := MinValIdx(row)

		// Iterate through the columns in this row
		c := []int{}
		// Loop over column indexes
		for r := range nRows {
			c = append(c, m[r][colIdx])
		}
		_, lMaxVal := MaxValIdx(c)

		if rMinVal == lMaxVal {
			out = append(out, lMaxVal)
		}
	}
	return out
}

func MinValIdx(arr []int) (int, int) {
	if len(arr) == 0 {
		return -1, -1
	}

	idx := 0
	min := arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			idx = i
			min = arr[i]
		}
	}

	return idx, min
}

func MaxValIdx(arr []int) (int, int) {
	if len(arr) == 0 {
		return -1, -1
	}

	idx := 0
	max := arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			idx = i
			max = arr[i]
		}
	}

	return idx, max
}

// A given some matrix of size M*N, a number is lucky if it is the minimum in it's row and the maximum in it's column
// We can't get around visiting every element, otherwise we won't know the minima/maxima, so the best we can do is going to be O(M*N)
// Now like other questions, we only need to keep the minima/maxima, which means we can discard values that are not
// So our strategy will be to iterate the matrix once, finding the minimum value of the row, and then check if this is the minimum value in the column too
// We'll only look at the values in the column where we found the minimum row value

// Here's an illustration to better understand the code
// [[r1c1, r1c2, r1c3, r1c4],
//  [r2c1, r2c2, r2c3, r2c4],
//  [r3c1, r3c2, r3c3, r3c4]]
// We will loop over the rows i.e. in our first loop we will get [r1c1, r1c2, r1c3, r1c4]
// Then we will find the minimum of this, suppose it is r1c3, and return the value r1c3 in addition to it's column index 2
// Then we will look at only column two, and calculate the max of that column
//                |
//                v
// [[r1c1, r1c2, r1c3, r1c4],
//  [r2c1, r2c2, r2c3, r2c4],
//  [r3c1, r3c2, r3c3, r3c4]]
//                ^
//                |
// The calculation of the column max is done by looping over the number of columns and adding the value for the column index to a temp array
// then taking this temp array and finding the maximum of it

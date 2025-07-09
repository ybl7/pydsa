package main

import (
	"fmt"
	"testing"
)

func TestRottenOranges(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name: "Basic case - all oranges rot",
			grid: [][]int{
				{2, 1, 1},
				{1, 1, 0},
				{0, 1, 1},
			},
			expected: 4,
		},
		{
			name: "Impossible case - isolated fresh orange",
			grid: [][]int{
				{2, 1, 1},
				{0, 1, 1},
				{1, 0, 1},
			},
			expected: -1,
		},
		{
			name: "No fresh oranges - all already rotten",
			grid: [][]int{
				{2, 2, 2},
				{2, 2, 2},
				{2, 2, 2},
			},
			expected: 0,
		},
		{
			name: "No oranges at all - empty grid",
			grid: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			expected: 0,
		},
		{
			name: "Single fresh orange, no rotten",
			grid: [][]int{
				{1},
			},
			expected: -1,
		},
		{
			name: "Single rotten orange",
			grid: [][]int{
				{2},
			},
			expected: 0,
		},
		{
			name: "Mixed with empty spaces",
			grid: [][]int{
				{2, 1, 0, 0, 1},
				{1, 0, 1, 2, 1},
				{1, 0, 0, 2, 1},
			},
			expected: 2,
		},
		{
			name: "Large grid - linear spread",
			grid: [][]int{
				{2, 1, 1, 1, 1},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
			},
			expected: 4,
		},
		{
			name: "Multiple rotten sources",
			grid: [][]int{
				{2, 1, 1, 1, 2},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
			},
			expected: 4,
		},
		{
			name: "Blocked path",
			grid: [][]int{
				{2, 1, 0, 1},
				{1, 0, 0, 1},
				{1, 0, 0, 1},
			},
			expected: -1,
		},
		{
			name: "Corners case",
			grid: [][]int{
				{2, 0, 1},
				{0, 0, 0},
				{1, 0, 2},
			},
			expected: -1,
		},
		{
			name: "Single row",
			grid: [][]int{
				{2, 1, 1, 1, 1},
			},
			expected: 4,
		},
		{
			name: "Single column",
			grid: [][]int{
				{2},
				{1},
				{1},
				{1},
			},
			expected: 3,
		},
		{
			name: "Spiral pattern",
			grid: [][]int{
				{1, 1, 1, 1, 1},
				{1, 0, 0, 0, 1},
				{1, 0, 2, 0, 1},
				{1, 0, 0, 0, 1},
				{1, 1, 1, 1, 1},
			},
			expected: -1,
		},
		{
			name: "Immediate neighbors only",
			grid: [][]int{
				{0, 2, 0},
				{1, 1, 1},
				{0, 1, 0},
			},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of the grid since the function modifies it
			gridCopy := make([][]int, len(tt.grid))
			for i := range tt.grid {
				gridCopy[i] = make([]int, len(tt.grid[i]))
				copy(gridCopy[i], tt.grid[i])
			}

			result := RottenOranges(gridCopy)
			if result != tt.expected {
				t.Errorf("Test %s failed: expected %d, got %d", tt.name, tt.expected, result)
				fmt.Printf("Grid:\n")
				for _, row := range tt.grid {
					fmt.Printf("%v\n", row)
				}
				fmt.Printf("Expected: %d, Got: %d\n\n", tt.expected, result)
			}
		})
	}
}

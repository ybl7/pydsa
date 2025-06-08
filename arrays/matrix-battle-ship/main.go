package main

import "fmt"

func main() {
	tests := []struct {
		name     string
		board    [][]string
		expected int
	}{
		{
			name: "Two battleships - one horizontal, one vertical",
			board: [][]string{
				{"X", ".", ".", "X"},
				{".", ".", ".", "X"},
				{".", ".", ".", "X"},
			},
			expected: 2,
		},
		{
			name: "Single horizontal battleship",
			board: [][]string{
				{"X", "X", "X"},
			},
			expected: 1,
		},
		{
			name: "Single vertical battleship",
			board: [][]string{
				{"X"},
				{"X"},
				{"X"},
			},
			expected: 1,
		},
		{
			name: "No battleships - all empty",
			board: [][]string{
				{".", ".", "."},
				{".", ".", "."},
			},
			expected: 0,
		},
		{
			name: "Multiple single cell battleships",
			board: [][]string{
				{"X", ".", "X"},
				{".", ".", "."},
				{"X", ".", "X"},
			},
			expected: 4,
		},
		{
			name: "Complex layout with mixed battleships",
			board: [][]string{
				{"X", "X", ".", "X"},
				{".", ".", ".", "X"},
				{".", "X", "X", "."},
			},
			expected: 3,
		},
		{
			name: "Single cell battleship",
			board: [][]string{
				{"X"},
			},
			expected: 1,
		},
		{
			name: "Alternating single battleships",
			board: [][]string{
				{"X", ".", "X", ".", "X"},
			},
			expected: 3,
		},
		{
			name:     "Empty board",
			board:    [][]string{},
			expected: 0,
		},
		{
			name: "Two separate horizontal battleships",
			board: [][]string{
				{"X", "X", ".", "X", "X"},
			},
			expected: 2,
		},
	}

	passedTests := 0
	totalTests := len(tests)

	fmt.Println("Running Battleships in a Board Tests")
	fmt.Println("====================================")

	for _, test := range tests {
		result := MatBattleShip(test.board)
		passed := result == test.expected

		if passed {
			fmt.Printf("✓ PASS: %s\n", test.name)
			passedTests++
		} else {
			fmt.Printf("✗ FAIL: %s\n", test.name)
			fmt.Printf("  Board:\n")
			for _, row := range test.board {
				fmt.Printf("    %v\n", row)
			}
			fmt.Printf("  Expected: %d\n", test.expected)
			fmt.Printf("  Got:      %d\n", result)
			fmt.Println()
		}
	}

	fmt.Println("====================================")
	fmt.Printf("Tests passed: %d/%d\n", passedTests, totalTests)

	if passedTests == totalTests {
		fmt.Println("🎉 All tests passed!")
	} else {
		fmt.Printf("❌ %d tests failed\n", totalTests-passedTests)
	}
}

// Helper function to visualize the board (for debugging)
func printBoard(board [][]string) {
	fmt.Println("Board visualization:")
	for _, row := range board {
		for _, cell := range row {
			fmt.Print(cell + " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func MatBattleShip(b [][]string) int {
	// Cases where there is nothing to search, i.e. 0 battleships without a board
	if len(b) == 0 {
		return 0
	}
	if len(b[0]) == 0 {
		return 0
	}
	count := 0
	j := len(b)
	i := len(b[0])

	for x := 0; x < i; x++ {
		for y := 0; y < j; y++ {
			if b[y][x] == "." {
				continue
			}
			if b[y][x] == "X" {
				emptyLeft, emptyTop := false, false
				// Check left
				if x-1 < 0 || b[y][x-1] == "." {
					emptyLeft = true
				}
				// Check top
				if y-1 < 0 || b[y-1][x] == "." {
					emptyTop = true
				}
				if emptyLeft && emptyTop {
					count++
				}
			}
		}
	}
	return count
}

// This question is quite tricky if you don't know what to look for, the key lies in the hint that is given: that no two battleships are adjacent
// Rephrased, this means wherever you find a battleship, you are guaranteed to find only empty squares around it.
// If you do happen to find some squares that are together, then these must belong to the same battleship and should only be counted once together.
// So it suffices to find a single square to identify a battleship.

// One thing to note is that battle ships are either 1 * k or k * 1 in shape.
// So they are either horizonal or vertical lines, they can't take on more exotic shapes.
// This is useful since it tells us that a battleship square will have at most two neighbours, but it also tells us how to identiy them.

// Regarldess of whether you are k * 1 or 1 * k shape, there will always be one square, the first square, that is top most and leftmost.
// So if we can just identify this, then we've identified a battle ship. The key insight is that the ship will ALWAYS have no TOP or LEFT neighbours.
// So these squares are the ones we look for that we look for.
// This necessarily implies that all other battleship squares WILL have at least a top or a right neighbour.

// We could also use the bottom most and right most square with the condition that the square will have no right or bottom neighbours.
// But this is an unnatural choice since we iterate matrices from left to right, top to bottom.

// So our strategy will be to loop over the matrix once, and if we find an X that has no left or top neighbours, we increment a counter.
// Our time complexity can be no better than O(M*N) since we must visit every node at least once, as there is no way to know if an X is a new ship
// or if it belongs to an old ship without checking it. For each X we will perform at most 2 calculations to check if it's a ship head.
// We'll perform lookups against it's top and left neighbour, both can happen in O(1) since looking up elements in arrays by index is an O(1) operation.

// Now I think that Meta may ask an extended version of this question:

// Given a board return the orientation of the battleship.
// Then return, all the coordinates with a battleship given size of ship.
// Follow up: write a function to randomly place a battleship of size n on board.

// We can extend our solution quite simply. Once we find the head of the battleship (we modify the above code to return the coordinates), i.e. the left most.
// We call a helper function to check if it's a horizontal or vertical battleship by checking if it has a botton or right neighbour.
// Once we get the orientation. We call another helper function that takes a current (x,y) and checks if the next (x+1.y) or (x,y+1) is also there.
// Now apparently we might be able to assume ships of constant length, so this second function is not necessary.
// Once we get the start coordinates and orientation, we should easily imply the other coordinates.

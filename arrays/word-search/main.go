package main

import "fmt"

func main() {
	fmt.Println("=== Word Search Test Cases ===\n")

	// Test Case 1: Basic Valid Path
	fmt.Println("Test Case 1: Basic Valid Path")
	board1 := [][]string{
		{"A", "B", "C", "E"},
		{"S", "F", "C", "S"},
		{"A", "D", "E", "E"},
	}
	word1 := "ABCCED"
	result1 := WordSearch(board1, word1)
	expected1 := true
	fmt.Printf("Board: %v\n", board1)
	fmt.Printf("Word: %s\n", word1)
	fmt.Printf("Result: %t, Expected: %t", result1, expected1)
	if result1 == expected1 {
		fmt.Println(" ✓ PASS")
	} else {
		fmt.Println(" ✗ FAIL")
	}
	fmt.Println()

	// Test Case 2: Word Not Found
	fmt.Println("Test Case 2: Word Not Found")
	board2 := [][]string{
		{"A", "B", "C", "E"},
		{"S", "F", "C", "S"},
		{"A", "D", "E", "E"},
	}
	word2 := "ABCB"
	result2 := WordSearch(board2, word2)
	expected2 := false
	fmt.Printf("Board: %v\n", board2)
	fmt.Printf("Word: %s\n", word2)
	fmt.Printf("Result: %t, Expected: %t", result2, expected2)
	if result2 == expected2 {
		fmt.Println(" ✓ PASS")
	} else {
		fmt.Println(" ✗ FAIL")
	}
	fmt.Println()

	// Test Case 3: Single Character
	fmt.Println("Test Case 3: Single Character")
	board3 := [][]string{{"A"}}
	word3 := "A"
	result3 := WordSearch(board3, word3)
	expected3 := true
	fmt.Printf("Board: %v\n", board3)
	fmt.Printf("Word: %s\n", word3)
	fmt.Printf("Result: %t, Expected: %t", result3, expected3)
	if result3 == expected3 {
		fmt.Println(" ✓ PASS")
	} else {
		fmt.Println(" ✗ FAIL")
	}
	fmt.Println()

	// Test Case 4: Word Requires Backtracking
	fmt.Println("Test Case 4: Word Requires Backtracking")
	board4 := [][]string{
		{"A", "B", "C", "E"},
		{"S", "F", "E", "S"},
		{"A", "D", "E", "E"},
	}
	word4 := "ABCESEEEFS"
	result4 := WordSearch(board4, word4)
	expected4 := true
	fmt.Printf("Board: %v\n", board4)
	fmt.Printf("Word: %s\n", word4)
	fmt.Printf("Result: %t, Expected: %t", result4, expected4)
	if result4 == expected4 {
		fmt.Println(" ✓ PASS")
	} else {
		fmt.Println(" ✗ FAIL")
	}
	fmt.Println()

	// Test Case 5: Large Grid with Long Word
	fmt.Println("Test Case 5: Large Grid - Word Not Possible")
	board5 := [][]string{
		{"A", "A", "A", "A", "A", "A"},
		{"A", "A", "A", "A", "A", "A"},
		{"A", "A", "A", "H", "O", "L"},
		{"A", "A", "A", "E", "L", "L"},
		{"A", "A", "A", "A", "O", "O"},
	}
	word5 := "HELLOWORLD"
	result5 := WordSearch(board5, word5)
	expected5 := false
	fmt.Printf("Board: %v\n", board5)
	fmt.Printf("Word: %s\n", word5)
	fmt.Printf("Result: %t, Expected: %t", result5, expected5)
	if result5 == expected5 {
		fmt.Println(" ✓ PASS")
	} else {
		fmt.Println(" ✗ FAIL")
	}
	fmt.Println()

	// Test Case 6: Spiral Pattern
	fmt.Println("Test Case 6: Spiral Pattern")
	board6 := [][]string{
		{"S", "P", "I", "R"},
		{"A", "A", "A", "A"},
		{"L", "A", "A", "L"},
		{"S", "P", "I", "R"},
	}
	word6 := "SPIRAL"
	result6 := WordSearch(board6, word6)
	expected6 := true
	fmt.Printf("Board: %v\n", board6)
	fmt.Printf("Word: %s\n", word6)
	fmt.Printf("Result: %t, Expected: %t", result6, expected6)
	if result6 == expected6 {
		fmt.Println(" ✓ PASS")
	} else {
		fmt.Println(" ✗ FAIL")
	}
	fmt.Println()

	// Summary
	fmt.Println("=== Test Summary ===")
	fmt.Println("Replace the placeholder WordSearch function with your implementation")
	fmt.Println("Expected results:")
	fmt.Println("- Test 1: true  (ABCCED exists)")
	fmt.Println("- Test 2: false (ABCB - can't reuse B)")
	fmt.Println("- Test 3: true  (Single A matches)")
	fmt.Println("- Test 4: true  (ABCESEEEFS with backtracking)")
	fmt.Println("- Test 5: false (HELLOWORLD not possible)")
	fmt.Println("- Test 6: true  (SPIRAL in spiral pattern)")
}

func WordSearch(mat [][]string, word string) bool {
	for y, m := range mat {
		for x, n := range m {
			// If we find our start string call RecSearch
			if n == word[:1] {
				return RecSearch(mat, x, y, word)
			}
		}
	}
	return false
}

// Recursively search neigbours
func RecSearch(mat [][]string, x, y int, word string) bool {
	xMax := len(mat[0])
	yMax := len(mat)

	if len(word) < 1 {
		return true
	}
	char := word[:1]
	rest := word[1:]

	// One way to check that we are on the board is to check before the recursive call where we are currently and decide what the adjacent cells should be
	// The better way to do it is to just make a recursive call with all 4 neighbour coords, and then have a validation step at the start of the recursive call to reject invalid indices
	if x < 0 || x >= xMax || y < 0 || y >= yMax {
		return false
	}

	// This entry is not the letter we need
	if mat[y][x] != char {
		return false
	}

	// The question says that the same cell may NOT be used more than once, this is important since our search algorithm may oscillate between two cells, suppose I have two adjacent cels N and O
	// If my target word is nonononononono for example, if I can reuse cells then the alforithm will tell me that I can construct nonononononono despite not having enough cells, so we need to block this cell as used for future recursive calls
	// But not to worry, we pass the board in by value, which means that calls made to search the board that aren't the children of this call will not have that cell blocked. If we get to this stage it means we have found the next valid char.
	mat[y][x] = "X"

	// General cases
	return RecSearch(mat, x, y-1, rest) || RecSearch(mat, x, y+1, rest) || RecSearch(mat, x-1, y, rest) || RecSearch(mat, x+1, y, rest)
}

// Given an m x n grid of characters board and a string word, return true if word exists in the grid.
// The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once.

// Approach
// Traverse matrix and look for start character of work
// If found, recursively look for next character in word
// Each time we decrement the string we are passing into the function as we find letters

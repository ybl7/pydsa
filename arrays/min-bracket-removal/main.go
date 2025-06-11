package main

import (
	"fmt"

	stacks "github.com/idsulik/go-collections/v3/stack/arraystack"
)

func main() {
	// Test cases for MinBracketRemoval function
	testCases := []struct {
		input    string
		expected int
		desc     string
	}{
		// Basic cases
		{"", 0, "Empty string"},
		{"()", 0, "Single valid pair"},
		{"(()", 1, "One unmatched opening"},
		{"())", 1, "One unmatched closing"},

		// Simple invalid cases
		{"(", 1, "Single opening bracket"},
		{")", 1, "Single closing bracket"},
		{"((", 2, "Two opening brackets"},
		{"))", 2, "Two closing brackets"},

		// Mixed valid and invalid
		{"()())", 1, "Valid pairs with extra closing"},
		{"((()", 2, "Multiple opening with one closing"},
		{"((()))", 0, "Nested valid brackets"},
		{"()(()", 1, "Valid pair followed by unmatched"},

		// Complex cases
		{"((())", 1, "Nested with one extra opening"},
		{"()()())", 1, "Multiple valid pairs with extra closing"},
		{"(((((", 5, "All opening brackets"},
		{")))))", 5, "All closing brackets"},

		// Alternating patterns
		{")(", 2, "Closing then opening"},
		{")()(", 2, "Alternating invalid pattern"},
		{"())()", 1, "Valid-invalid-valid pattern"},

		// Larger test cases
		{"((()())", 1, "Nested and sequential mix"},
		{"(()(()", 2, "Complex nested pattern"},
		{"())(()", 2, "Mixed invalid positions"},
		{"((())())", 0, "Complex valid nested pattern"},

		// Edge cases with letters (if problem allows)
		{"(a)", 0, "Single letter in valid brackets"},
		{"(a))", 1, "Letter with extra closing"},
		{"((a)", 1, "Letter with extra opening"},

		// Stress test cases
		{"((((((((()", 8, "Many unmatched opening"},
		{"))))))))))", 10, "Many unmatched closing"},
		{"()()()()()", 0, "Many valid pairs"},
		{"()()()()((", 2, "Valid pairs ending with unmatched"},
	}

	fmt.Println("Running MinBracketRemoval test cases...")
	fmt.Println("=" + fmt.Sprintf("%50s", "="))

	passed := 0
	total := len(testCases)

	for i, tc := range testCases {
		result := MinBracketRemoval(tc.input)
		status := "PASS"
		if result != tc.expected {
			status = "FAIL"
		} else {
			passed++
		}

		fmt.Printf("Test %2d: %-25s | Input: %-15s | Expected: %2d | Got: %2d | %s\n",
			i+1, tc.desc, fmt.Sprintf("\"%s\"", tc.input), tc.expected, result, status)

		fmt.Println()
	}

	fmt.Println("=" + fmt.Sprintf("%50s", "="))
	fmt.Printf("Results: %d/%d tests passed (%.1f%%)\n", passed, total, float64(passed)/float64(total)*100)

	if passed == total {
		fmt.Println("🎉 All tests passed!")
	} else {
		fmt.Printf("❌ %d test(s) failed\n", total-passed)
	}
}

func MinBracketRemoval(s string) int {
	st := stacks.New[rune](len(s))

	// iterating a string yields bytes, will assume chars are asci
	for _, c := range s {
		top, _ := st.Peek()
		if c == '(' {
			st.Push('(')
			// We use the same stack to store both opening and closing, we must make the destinction that we only close (pop) when there is a bracket open
		} else if c == ')' && !st.IsEmpty() && top == '(' {
			st.Pop()
		} else if c == ')' {
			st.Push(')')
		}
		fmt.Printf("st: %v\n", st)
	}
	return st.Len()
}

// The idea here is very simple. When are brackets invalid?
// we traverse strings from left to right so invalid brackets occour when:
// When there are more brackets opened than closed in total
// When a there are more closing brackets before opening brackets
// So all we need to do is count these occurances, but violations that happen earlier, may get fixed later
// e.g. (((...................................)..........................).......)
// So we need to keep a running sum of the opening brackets, only if at the end we have more open overall can we conclude that they need to be removed
// If ever we have more closing brakcets than opening, we know immediately that these should be removed.
// So we have a counter that is constantly increasing and decreasing, and we want to take it's value right at the end of the interation to see how many violations
// This is a classic stack problem, we push and pop on the stack as necessary and take it's length at the end

package main

import (
	"fmt"
	"strconv"

	stacks "github.com/idsulik/go-collections/v3/stack/arraystack"
)

// Test cases for Basic Calculator II
func runTests() {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Basic multiplication precedence",
			input:    "3+2*2",
			expected: 7,
		},
		{
			name:     "Division with truncation",
			input:    " 3/2 ",
			expected: 1,
		},
		{
			name:     "Mixed operations with spaces",
			input:    " 3+5 / 2 ",
			expected: 5,
		},
		{
			name:     "Complex precedence",
			input:    "2*3-1+4/2",
			expected: 7,
		},
		{
			name:     "Division truncation",
			input:    "7/3",
			expected: 2,
		},
		{
			name:     "Multiple operations",
			input:    "1+2*3+4*5-6/2",
			expected: 24,
		},
		{
			name:     "Single number",
			input:    "42",
			expected: 42,
		},
		{
			name:     "Zero handling",
			input:    "0+1*2-0/1",
			expected: 2,
		},
		{
			name:     "Large numbers",
			input:    "100-50+25*2/5",
			expected: 60,
		},
		{
			name:     "Sequential operations",
			input:    "8*2/4*3/2",
			expected: 6,
		},
	}

	passed := 0
	total := len(testCases)

	for _, tc := range testCases {
		result := BasicCalc(tc.input)
		if result == tc.expected {
			fmt.Printf("✓ PASS: %s\n", tc.name)
			passed++
		} else {
			fmt.Printf("✗ FAIL: %s\n", tc.name)
			fmt.Printf("  Input:    \"%s\"\n", tc.input)
			fmt.Printf("  Expected: %d\n", tc.expected)
			fmt.Printf("  Got:      %d\n", result)
			fmt.Println()
		}
	}

	fmt.Printf("\nResults: %d/%d tests passed\n", passed, total)
	if passed == total {
		fmt.Println("🎉 All tests passed!")
	}
}

func main() {
	fmt.Println("Running Basic Calculator II Tests")
	fmt.Println("==================================")

	runTests()
}

func BasicCalc(s string) int {
	stkO := stacks.New[rune](100)
	stkN := stacks.New[int](100)
	curr := 0
	for i, e := range s {
		if e == ' ' {
			// skip spaces
			continue
		}
		if '0' <= e && e <= '9' {
			digit, _ := strconv.ParseInt(s[i:i+1], 10, 64)
			curr = curr*10 + int(digit)
		}
		if e == '+' || e == '-' || e == '*' || e == '/' {
			// we've hit an operator so push whatever operand we have onto the queue then clear it
			stkN.Push(curr)
			curr = 0

			// then push onto the stack based on the logic discussed below
			if stkO.IsEmpty() {
				stkO.Push(e)
			} else {
				top, _ := stkO.Peek()
				// if the current operator at top has higher prio than the incoming one then it must be done first
				if PrioAoverB(top, e) {
					t, _ := stkO.Pop()
					b, _ := stkN.Pop()
					a, _ := stkN.Pop()
					v := Operate(a, b, t)
					stkN.Push(v)
					// Don't forget to push the current operator on after we have evaluated the expression
					stkO.Push(e)
				} else {
					stkO.Push(e)
				}
			}
		}
	}
	// push on last number
	if curr != 0 {
		stkN.Push(curr)
	}

	// After reaching the end of the string we need to finish the calculation
	for !stkO.IsEmpty() {
		t, _ := stkO.Pop()
		b, _ := stkN.Pop()
		a, _ := stkN.Pop()
		v := Operate(a, b, t)
		stkN.Push(v)
	}
	out, _ := stkN.Pop()
	return out
}

func PrioAoverB(a, b rune) bool {
	prio := map[rune]int{
		'+': 0,
		'-': 0,
		'/': 1,
		'*': 1,
	}
	// the >= is very important, if operation are the same prio, we do the first one, this is crucial
	// since when operations have the same prio, we have to do them from left to right i.e. 6-1+4 is (6-1)+4 NOT 6-(1+4)
	return prio[a] >= prio[b]
}

func Operate(a, b int, c rune) int {
	var out int

	switch c {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	case '/':
		return a / b
	}

	return out
}

// Given a string s which represents an expression, evaluate this expression and return its value.
// The integer division should truncate toward zero.
// You may assume that the given expression is always valid. All intermediate results will be in the range of [-231, 231 - 1].
// Note: You are not allowed to use any built-in function which evaluates strings as mathematical expressions, such as eval().
// 1 <= s.length <= 3 * 105
// s consists of integers and operators ('+', '-', '*', '/') separated by some number of spaces.
// s represents a valid expression.
// All the integers in the expression are non-negative integers in the range [0, 231 - 1].
// The answer is guaranteed to fit in a 32-bit integer.

// Suppose we have some expression like a op1 b op2 c, where a, b, and c are numbers and op1, and op2 are operators
// Let's not some obvious facts that we might gloss over day to day
// 1) operators operate on 2 operands (in this question at least)
// 2) numbers involved in the same operation can at most be separated by one operator, i.e. numbers involved in the #
//	  same operation are adjacent, you can't have a being operated with c, as in our example above
// 3) operators have a different priority, but like numbers, you only need to consider any two at a time, either op1
//	  has greater priority than op2 or the other way round, there is no need to get any other operators involved
// 4) to know the correct order of operations for some number between two operators we need to see both operators
// 5) to know the result of an operation between two numbers we need to know both numbers
// These - I'll call them adjacency properties - mean that in at any given point we are only really looking at two adjacent
// number (separated by an operator) or two adjacent operators separated by numbers
// We aren't working with prefix or postfix notation where it's possible to have multiple adjacent numbers or operators

// So with that in mind, whenever we operate on at most two adjacent things and we need to know them and be able access them,
// a convenient data structure to use is a stack. We might need to reverse the order in which we do operations
// (relative to the order they show up in a string) so we use a stack for the operations

// Let's take an example to make this more clear a op2 b op1 c op3 d, and let's say prio(op3) > prio(op1) > prio(op2)
// We will create two stacks and iterate down the string, stack 1 for numbers, stack 2 for operators

// stack_nums = [a]
// stack_ops = []
// the first entity in the string is a number, push it onto stack nums

// stack_nums = [a]
// stack_ops = [op2]
// the second entity is an operator, push it onto stack 2

// stack_nums = [a, b]
// stacks_ops = [op2]
// push b onto the stack now we have enough entities to do a op1 b, but this might not be correct,
// since the b is involved in another operation that could be of higher precedence, so before we push this next operator
// onto the stack we need to check if it's higher or lower priority than our current operator at the top of the stack

// if prio(op2) > prio(op1) then we need to do op2 before op1, so we pop off the stack and do a op2 b,
// once we get this result we push it back onto the nums stack, since this composite number will be used in later calcs
// now we also push op1 onto the stack
// stack_nums = [a op2 b]
// stacks_ops = [op1]

// Then we just continue and push c onto stack_nums, and then attempt to push op3 onto stack_ops, but once again we must
// check prio(op1) > prio(op3), if it is we must do the op1 operation now, as it turns out it isn't so we push op3 on
// stack_nums = [a op2 b, c]
// stacks_ops = [op1, op3]

// then we push d onto stack nums, and we've hit the end of string
// stack_nums = [a op2 b, c, d]
// stacks_ops = [op1, op3]
// at which point we pop off stack_nums and stack_ops until both stacks are empty, so we do c op3 d

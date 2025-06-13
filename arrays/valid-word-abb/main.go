package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	testCases := []TestCase{
		// Valid cases
		{"internationalization", "i12iz4n", true, "Classic example - mixed chars and numbers"},
		{"word", "4", true, "Entire word abbreviated as number"},
		{"apple", "apple", true, "No abbreviation - exact match"},
		{"substitution", "s10n", true, "Number at end"},
		{"hello", "h3o", true, "Number in middle"},
		{"programming", "p2g2mming", true, "Multiple numbers"},
		{"localization", "l10n", true, "Common l10n abbreviation"},
		{"", "", true, "Both empty strings"},
		{"a", "1", true, "Single char abbreviated"},
		{"substitution", "sub4u4", true, "ends in abbreviation"},

		// Invalid cases - character mismatches
		{"apple", "a2f", false, "Character mismatch at end"},
		{"test", "best", false, "First character mismatch"},

		// Invalid cases - leading zeros
		{"hi", "01", false, "Leading zero - two digits"},
		{"hello", "h03o", false, "Leading zero in middle"},
		{"word", "w0rd", false, "Zero as leading digit"},

		// Invalid cases - length/count mismatches
		{"hello", "h10", false, "Skip count too large"},
		{"", "1", false, "Empty word, non-empty abbr"},
	}

	fmt.Println("Running Valid Word Abbreviation Tests...")
	fmt.Println("=====================================")

	passed := 0
	failed := 0

	for i, tc := range testCases {
		// result := validWordAbbreviation(tc.full, tc.abbr)
		result := ValidWordAbb(tc.full, tc.abbr)
		status := "PASS"
		if result != tc.expected {
			status = "FAIL"
			failed++
		} else {
			passed++
		}

		fmt.Printf("Test %2d: %s\n", i+1, status)
		fmt.Printf("  Word: '%s'\n", tc.full)
		fmt.Printf("  Abbr: '%s'\n", tc.abbr)
		fmt.Printf("  Expected: %t, Got: %t\n", tc.expected, result)
		fmt.Printf("  Description: %s\n", tc.desc)
		fmt.Println()
	}
}

type TestCase struct {
	full     string
	abbr     string
	expected bool
	desc     string
}

// My first solution
func ValidWordAbb(full, abbr string) bool {
	i, j, k := 0, 0, 0

	for ; j < len(abbr); j++ {
		// Check if j is number by converting byte to rune and compare since runes are just int32s
		if '0' <= rune(abbr[j]) && rune(abbr[j]) <= '9' {
			if rune(abbr[j]) == '0' && k == 0 {
				return false
			}
			// This is a nifty way to get an int from a number in string form e.g. 125 = 1 * 10 * 10 + 2 * 10 + 5
			v, err := strconv.ParseInt(abbr[j:j+1], 10, 64)
			HandleError(err)
			k = k*10 + int(v)

			// If abbr ends in j and the value under or over shoots return false
			// I always let j run to the end, so I pin it to be the correct length, meaning i will either under or overshoot, right? i < j == j > i and i > j == j < i; I'm just saying what the "ideal" solution
			// is doing in a different way
			if j+1 >= len(abbr) {
				i = i + k
				if i > len(full) || i < len(full) {
					return false
				}
			}
		} else {
			// The if here is not necessary, if k == 0 jumping x by k does nothing, and if it's not, we need to jump and reset, we duplicate the check abbr[j] != full[i] also
			if k == 0 {
				if abbr[j] != full[i] {
					return false
				}
			} else {
				i = i + k
				if abbr[j] != full[i] {
					return false
				}
				k = 0
			}
			i++
		}
	}
	return true
}

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// A solutionn off the web, does it pass as many test cases as mine? Yes it's shorter and better, but harder to understand
// func validWordAbbreviation(word string, abbr string) bool {
// 	m, n := len(word), len(abbr)
// 	i, j, x := 0, 0, 0
// 	for ; i < m && j < n; j++ {
// 		if abbr[j] >= '0' && abbr[j] <= '9' {
// 			if x == 0 && abbr[j] == '0' {
// 				return false
// 			}
// 			x = x*10 + int(abbr[j]-'0')
// 		} else {
// 			i += x
// 			x = 0
// 			if i >= m || word[i] != abbr[j] {
// 				return false
// 			}
// 			i++
// 		}
// 	}
// 	return i+x == m && j == n
// }

// I didn't think about this well enough initially and went into "proof by example" mode, let's state some facts
// An abbreviated string should be the same length as the original string when you take into account the abbreviations
// By the end of the algorithm, if the strings do indeed match, then both pointers i and j must be at value len of the respective string (by definition)
// Since I choose to iterate down j, this is already guaranteed, there is no way for j to over or undershoot this len(abbr)
// If j hits a number, we pause i until we get the full number that we need to jump
// If j hits a character, we check if we need to jump i and do so, (jumping i by 0 does no harm), then we reset the jump (we have to do this by definition since we aren't in a number any more)
// Then we check if the values at i and j are equal in their respective strings, if not we fail, we must also check that we don't accidentally jump too much, then we increment both i and j to move onto the next char
// If at the end, we reach the end of j, we either have a jump if a number was the last thing in abbr, or we have zero jump, in both cases we should jump (it's a noop in the 0 case)
// Then we check, after this final jump if i = len(full), the code from the internet also checks if j == len(abbr). Now the loop they use also checks i < m, so this is the reason why j can sometimes never reach
// len(abbr), in the rare cases where i exits out before it (for example if there is a huge number like 100 in abbr, i will always exit first). So let's work through the example again.

// i = 0, "internationalization"
//		   ^
// j = 0, "i12iz4n"
//		   ^
// k = 0
// if j is not a number, check if we need to jump i (covers the case when we just came out of a substring with a number), jump i if necessary, check if abbr[j] == full[i], and check if i is out of bounds
// j++, i++

// i = 1, "internationalization"
//		    ^
// j = 1, "i12iz4n"
//		    ^
// k = 0
// j is number so k = 1, j++

// i = 1, "internationalization"
//		    ^
// j = 2, "i12iz4n"
//		     ^
// k = 0
// j is number so k = 12, j++

// i = 1, "internationalization"
//		    ^
// j = 3, "i12iz4n"
//		      ^
// k = 0
// j is not a number, so we check if i needs to jump (it does by 12), so jump it (but then check we didn't accidentally OVERJUMP), i = 1 + 12 = 13
// Then check full[13] == abbr[3], if it's not we exit, then increment both i and j

// And that's really it, but we know by the end, if abbr really is an abbreviation of full, their pointers must both be the length of the respective strings
// This is the key insight that I was missing and it's why my original solution is so complex, well actually I can remove the if else in the second block (I don't really need to check if k = 0)
// adn taht length check from line 85 should be moved into the return statement and should check both arrays, so I guess I wasn't too far off (the k stuff is just syntax), and I think the exit condition
// that I chose is equivalent, since I always let j run to it's end, this always implies that i will NOT run to it's end if they are mismatched, i will either over or undershoot
// So actually, I don't think my solution is that bad, it looks worse for sure, but I think it achieves the exact same thing

// ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
// A string can be abbreviated by replacing any number of non-adjacent, non-empty substrings with their lengths.
// The lengths should not have leading zeros.

// We need to find if the distance in chars between two non digit endpoints is equal to the digit in between them.
// We will never have adjacent abbreviations so we are guaranteed to have two endpoints.
// If we do get an abbreviation like 55 as in the example, this should always be taken to mean 55 characters.
// So what should our approach be?
// Well we want to make sure that the abbreviation is actually equal to the number of characters.
// So we need to iterate down the abbreviated string to get the lengths of abbreviations, say 55.
// Then we need to iterate down the normal string and check if the number of characters matches those in the abbreviated string.
// Keeping in mind that we can't have leading 0s in the abbreviated string.

// This quesion is apparently a leetcode easy, but actually needs some thought
// The idea is to go through the abbreviated string and the unabbreviated string
// We'll have a pointer on each string. The j counter will iterate through the abbreviated string.
// Every time we hit a digit value we save it to some variable k, we do this until we reach a non digit value.
// Then we test if arr[i+k] = arr[j], that is if we start from i and skip k digits, the values in the abbreviated
// and full arrays should be equal, so i tracks j when j is a non digit, i.e. we increase i as normal
// but we jump i forward by whatever the value is in the abbreviated string and see if the next char in both strings are equal

// We need to go and check if the
// kubernetes can be abbreviated k8s
// abbr := "i12iz4n"
// full := "internationalization"

// i = 0, "internationalization"
//		   ^
// j = 0, "i12iz4n"
//		   ^
// k = 0
// if j is not a number and k == 0 and abbr[i] != full[i], return false
// else i++, j++

// i = 1, "internationalization"
//		    ^
// j = 1, "i12iz4n"
//		    ^
// k = 0
// if j is a number, stop incrementing i until we parse the number
// if k = 0 and j == 0 return false
// else parse k = 1
// j++

// i = 1, "internationalization"
//		    ^
// j = 2, "i12iz4n"
//		     ^
// k = 1
// if j is a number, stop incrementing i until we parse the number
// if k = 0 and j == 0 return false
// else parse k = 12
// j++

// i = 1, "internationalization"
//		    ^
// j = 3, "i12iz4n"
//		      ^
// k = 12
// if j is not a number and k != 0, i = i + k = 13
// if full[i] != abbr[j] return false else: i++ j++ k=0 (reset the jump)
// full[13] == i, abbr[3] == i so true

// i = 14, "internationalization"
//		                  ^
// j = 4, "i12iz4n"
//		       ^
// k = 0
// if j is not a number and k == 0 and abbr[i] != full[i], return false
// else i++ j++

// i = 15, "internationalization"
//		                   ^
// j = 5, "i12iz4n"
//		        ^
// k = 0
// if j is a number, stop incrementing i until we parse the number
// if k = 0 and j == 0 return false
// else parse k = 4
// j++

// i = 15, "internationalization"
//		                   ^
// j = 6, "i12iz4n"
//		         ^
// k = 4
// if j is a not a number and k != 0, i = i + k = 19
// if full[i] != abbr[j] return false else: i++ j++
// k=0 (reset the jump)
// full[19] == i, abbr[6] == i so true

// Return true at the end
// A bad case

// i = 0, "hello"
//		   ^
// j = 0, "h4"
//		   ^
// k = 0
// if j is not a number and k == 0 and abbr[i] != full[i], return false
// else i++, j++

// i = 1, "hello"
//		    ^
// j = 1, "h4"
//		    ^
// k = 0
// if j is a number, stop incrementing i until we parse the number
// if k = 0 and j == 0 return false
// else parse k = 4
// j++ but this would take us out of bounds, since we hit the end of abbr, we should do the check now
// if j+1 >= len(abbr) then, i = k + i = 5, i == len(full) return true, else return false

// So this is how we deal with the edge case where the abbr ends in a number

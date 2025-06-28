package main

import "math"

func MinWndSubstring(s, t string) string {
	out := ""
	minLen := math.MaxInt32

	want := map[rune]int{}
	have := map[rune]int{}
	valdCnt := len(t)
	currCnt := 0
	i := 0

	for _, e := range t {
		want[e]++
	}

	for j, c := range s {
		// Check if char is wanted, only count it if it is needed i.e. if want[c] > have[c]
		if _, ok := want[c]; ok && want[c] > have[c] {
			currCnt++
		}

		// Always add it to the current window state map
		have[c]++

		// Window contains enough characters
		for currCnt == valdCnt {
			if j+1-i < minLen {
				minLen = j + 1 - i
				out = s[i : j+1] // +1 since go excludes end interval, i.e. open endpoint
			}
			// Now shrink window until we violate validity
			r := []rune(s[i : i+1])[0] // get the string char at index i then convert to []rune that will contain only 1 rune, then take it's value using [0]
			have[r]--
			if _, ok := want[r]; ok && want[r] > have[r] {
				currCnt += -1
			}
			i++
		}
	}

	return out
}

// Of course we need to know how many chars to look for so we store them in a map, we need to iterate down t
// Key insight is to keep a map to store the counts of what we have seen, and to start shrinking windows once we have a valid window

// "ADOBECODEBANC" that contains all characters of "ABC"
// want == {A:1, B:1, C:1} and valdCnt == 3 (both never change)

// s == A, j == 0: currCnt == 1, have == {A:1},
// s == D, j == 1: currCnt == 1, have == {A:1, D:1}
// s == O, j == 2: currCnt == 1, have == {A:1, D:1, O:1}
// s == B, j == 3: currCnt == 2, have == {A:1, D:1, O:1, B:1}
// s == E, j == 4: currCnt == 2, have == {A:1, D:1, O:1, B:1, E:1}
// s == C, j == 5: currCnt == 3, have == {A:1, D:1, O:1, B:1, E:1, C:1}
//    currCnt == valdCnt == true, out == s[0:6] == "ADOBEC", r == A, have == {D:1, O:1, B:1, E:1, C:1}, currCnt == 2, i == 1
// ...

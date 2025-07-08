package main

import "fmt"

func main() {
	s1 := "banana"
	s2 := "chocolate"
	s3 := "abcdef"
	s4 := "delightful"

	fmt.Println(MaxPalindromicSubstr(s1))
	fmt.Println(MaxPalindromicSubstr(s2))
	fmt.Println(MaxPalindromicSubstr(s3))
	fmt.Println(MaxPalindromicSubstr(s4))
}

func MaxPalindromicSubstr(s string) string {
	dp := make([][]bool, len(s))
	for i, _ := range dp {
		dp[i] = make([]bool, len(s))
		for j, _ := range dp[i] {
			dp[i][j] = false
		}
	}

	// Set diagnol to true
	for i := range len(s) {
		dp[i][i] = true
	}

	max := 1
	var I, J int
	for i := len(dp) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			// If we find a cell, check if it extends a larger palindrom, i.e. check the southwest cell
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				if j-i+1 > max {
					max = j - i + 1
					I = i
					J = j
				}
			}
		}
	}

	return s[I : J+1]
}

//   a b c b d a
// a T F F F F T
// b   T F T F F
// c     T F F F
// b       T F F
// d		 T F
// a		   T

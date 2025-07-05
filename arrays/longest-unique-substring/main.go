package main

func MaxUniqSubstring(s string) int {
	max := 0

	track := map[string]bool{}
	i := 0

	for j := 0; j < len(s); j++ {
		for track[s[j:j+1]] {
			delete(track, s[i:i+1])
			i++
		}
		track[s[j:j+1]] = true

		if j-i+1 > max {
			max = j - i + 1
		}
	}

	return max
}

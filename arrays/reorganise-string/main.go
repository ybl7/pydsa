package main

import (
	"sort"
)

func ReorgString(str string) string {
	cnt := map[rune]int{}

	for _, r := range str {
		cnt[r]++
	}

	for _, freq := range cnt {
		if freq > (len(str)+1)/2 {
			return ""
		}
	}

	var keys []rune
	for k := range cnt {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return cnt[keys[i]] > cnt[keys[j]]
	})

	arr := make([]rune, len(str))
	i := 0

	for _, k := range keys {
		for cnt[k] > 0 {
			arr[i] = k
			cnt[k]--
			i += 2
			if i >= len(arr) {
				i = 1
			}
		}
	}

	return string(arr)
}

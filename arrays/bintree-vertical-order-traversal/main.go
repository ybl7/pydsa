package main

import (
	"fmt"
	"sort"
)

func main() {
	arr1 := []int{3, 9, 20, -101, -101, 15, 7}
	s1 := Solution(arr1)
	fmt.Println(s1)

	arr2 := []int{3, 9, 8, 4, 0, 1, 7}
	s2 := Solution(arr2)
	fmt.Println(s2)

	arr3 := []int{1, 2, 3, 4, 10, 9, 11, -101, 5, -101, -101, -101, -101, -101, -101, -101, -101, 6}
	s3 := Solution(arr3)
	fmt.Println(s3)
}

func TravAndPop(arr []int, i, x int, m map[int][]int) {
	// check if the current node exists, if not, return nothing
	if i >= len(arr) || arr[i] == -101 {
		return
	}

	if v, ok := m[x]; !ok {
		m[x] = []int{arr[i]}
	} else {
		v = append(v, arr[i])
		m[x] = v
	}

	// Left child
	TravAndPop(arr, 2*i+1, x-1, m)
	// Right child
	TravAndPop(arr, 2*i+2, x+1, m)
}

func Solution(arr []int) [][]int {
	m := make(map[int][]int)
	TravAndPop(arr, 0, 0, m)

	keys := []int{}
	for k, _ := range m {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	out := [][]int{}
	for _, k := range keys {
		out = append(out, m[k])
	}
	return out
}

// this q becomes easy when you think about it in cartesian coordinates
// for any value of x, you are trying to find all values of y
// so one approach is to iterate down the binary tree and just store all values by their x and y coordinate
// we can use a map for this, the keys of the map will be the x coordinates, and the values will be arrays of y coordinates
// the key to the question lies in this identification and in properly populating the map
// the leetcode question uses null to represent empty nodes in the array, go's zero value for ints is 0, can't use null
// so instead I'll use -101 since we are told the elements of the array must be between -100 and 100

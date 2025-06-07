package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(PartitionEqualSubarr([]int{1, 5, 11, 5}))               // true
	fmt.Println(PartitionEqualSubarr([]int{1, 2, 3, 5}))                // false
	fmt.Println(PartitionEqualSubarr([]int{4, 5, 6, 7, 8, 9, 1, 2, 3})) // true
	fmt.Println(PartitionEqualSubarr([]int{-3, 1, 2, 3, 4, 5, 6}))      // true
}

func PartitionEqualSubarr(arr []int) bool {
	sort.Ints(arr)

	sum := 0
	for _, e := range arr {
		sum += e
	}

	if sum%2 != 0 {
		return false
	}

	tgt := sum / 2
	// You must actually put something in the map or use make(), an empty map doesn't produce something that can be used
	store := map[int]bool{
		0: true,
	}

	for e := range arr {
		if e > tgt {
			break
		}

		for s, _ := range store {
			// If we get the tgt on this iteration return, otherwise store the new sum
			if s+e == tgt {
				return true
			}
			// Take all the current values in store and add e to them
			store[s+e] = true
		}
	}
	return false
}

// We want to be able to partition the array into two subsets, so the sum of each subset will naturally be half of the sum of the original array
// Given a general array [a,b,c,d], the brute force approach would be to calculate the sum of every subset
// Sum(a)
// Sum(a, b)
// Sum(a, c)
// Sum(a, d)
// Sum(a, b, c)
// Sum(a, b, d)
// Sum(a, c, d)
// Sum(a, b, c, d)
// Sum(b)
// Sum(b, c)
// Sum(b, d)
// Sum(b, c, d)
// Sum(c)
// Sum(c, d)
// Sum(d)
// So for our set of size 4 we have 15 possibilities (we exclude the empty set), generally the number of subsets for a set of size n is, 2^n
// Obviously a time complexity of 2^n is not acceptable, but we notice a few things, we don't need to calculate all of the sums since for example: Sum(a, b) = Sum(a) + Sum(b)

// Now we also need to calculate the total sum of the array. So we will do this in one pass, this is O(n), then we take the half of this number as this will be our target
// Now here is the clever part, if we need to find to subarrays that are equal, if we find a single element that is greater than the half sum, we can immediately conclude that this array doesn't satisfy the condition
// Also if the sum is not even, we exit immediately, since adding two equal values always produces an even value
// So assuming that we don't exit, this means that every element in the ordered array is necesarily < Sum(arr)/2
// For example in the case of [1, 5, 6, 10], our subsets are [1, 10] and [5, 6], but this demonstrates that we can't really make any conclusions regarding the indexes of the subsets

// So if we iterate down the sorted array. if Sum(arr)/2 = arr[j], then we are done, if it isn't, then we just need to remember the current value in case it become useful in later sums
// Recall we said that Sum(a, b) = Sum(a) + Sum(b), and similarly Sum(a, b, c) = Sum(a, b) + Sum(c) - so we can use this to save a little work

// Ideally what we want is to store every sum that we've seen until now, for example, once we visit element a in our example we want to store: {Sum(a)}
// Then once we visit b, we add to our store so we have {Sum(a), Sum(b), Sum(a, b)}, then once we visit c we add to our store again store {Sum(a), Sum(b), Sum(c), Sum(a, b), Sum(a, c), Sum(b, c), Sum(a, b, c)}
// Finally when we visit d our sum becomes: {Sum(a), Sum(b), Sum(c), Sum(d), Sum(a, b), Sum(a, c), Sum(a, d), Sum(b, c), Sum(b, d), Sum(c, d), Sum(a, b, d), Sum(a, c, d), Sum(a, b, c), Sum(b, c, d), Sum(a, b, c, d)}
// The pattern is quite clear now, we just take whatever we have in our store, then to get the updated store we take all the existed values and add the new element to them, then we add these to our store

// If we try and count the work that we do
// O(1) work for the check to see if our total sum is divisible by two
// O(nlog(n)) work to sort the array via a standard sorting algorithm like quicksort

// For the work inside the loop over the ordered array
// O(n) work initially as we loop through the array, we do at least O(1) work to check if we should break if the current value is > Sum(arr)/2
// But for each value of n, we actually do n iterations of a nested loop to update the store, now we could do a detailed analysis to see how this grows, I think it would be
// Sum(2^n) since the size of our store is 2^n, I don't see how this really helps us?
// Read this https://encyclopedia.pub/entry/28608

package main

import (
	"fmt"

	heap "github.com/idsulik/go-collections/v3/priorityqueue"
)

func main() {
	fmt.Println("Testing KthLargest function:")
	fmt.Println("==============================")

	// Test Case 1: Basic unsorted array
	nums1 := []int{3, 2, 1, 5, 6, 4}
	k1 := 2
	expected1 := 5
	result1 := KthLargest(nums1, k1)
	fmt.Printf("Test 1: nums=%v, k=%d\n", nums1, k1)
	fmt.Printf("Expected: %d, Got: %d", expected1, result1)
	if result1 == expected1 {
		fmt.Println(" ✓ PASS")
	} else {
		fmt.Println(" ✗ FAIL")
	}
	fmt.Println()

	// Test Case 2: Find maximum element (k=1)
	nums2 := []int{3, 2, 1, 5, 6, 4}
	k2 := 1
	expected2 := 6
	result2 := KthLargest(nums2, k2)
	fmt.Printf("Test 2: nums=%v, k=%d\n", nums2, k2)
	fmt.Printf("Expected: %d, Got: %d", expected2, result2)
	if result2 == expected2 {
		fmt.Println(" ✓ PASS")
	} else {
		fmt.Println(" ✗ FAIL")
	}
	fmt.Println()

	// Test Case 3: Find minimum element (k=length)
	nums3 := []int{7, 10, 4, 3, 20, 15}
	k3 := 6
	expected3 := 3
	result3 := KthLargest(nums3, k3)
	fmt.Printf("Test 3: nums=%v, k=%d\n", nums3, k3)
	fmt.Printf("Expected: %d, Got: %d", expected3, result3)
	if result3 == expected3 {
		fmt.Println(" ✓ PASS")
	} else {
		fmt.Println(" ✗ FAIL")
	}
	fmt.Println()

	// Test Case 4: Array with duplicates
	nums4 := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k4 := 4
	expected4 := 4
	result4 := KthLargest(nums4, k4)
	fmt.Printf("Test 4: nums=%v, k=%d\n", nums4, k4)
	fmt.Printf("Expected: %d, Got: %d", expected4, result4)
	if result4 == expected4 {
		fmt.Println(" ✓ PASS")
	} else {
		fmt.Println(" ✗ FAIL")
	}
	fmt.Println()

	// Test Case 5: Single element array
	nums5 := []int{1}
	k5 := 1
	expected5 := 1
	result5 := KthLargest(nums5, k5)
	fmt.Printf("Test 5: nums=%v, k=%d\n", nums5, k5)
	fmt.Printf("Expected: %d, Got: %d", expected5, result5)
	if result5 == expected5 {
		fmt.Println(" ✓ PASS")
	} else {
		fmt.Println(" ✗ FAIL")
	}
	fmt.Println()

	// Test Case 6: Negative numbers
	nums6 := []int{-1, -3, -2, -5, -4}
	k6 := 2
	expected6 := -2
	result6 := KthLargest(nums6, k6)
	fmt.Printf("Test 6: nums=%v, k=%d\n", nums6, k6)
	fmt.Printf("Expected: %d, Got: %d", expected6, result6)
	if result6 == expected6 {
		fmt.Println(" ✓ PASS")
	} else {
		fmt.Println(" ✗ FAIL")
	}
	fmt.Println()

	fmt.Println("Testing completed!")
}
func KthLargest(arr []int, k int) int {
	// if less(i, j) == true then i is higher priority, this is what we want since we have a minHeap
	h := heap.New(func(i, j int) bool {
		return i < j
	})

	for _, e := range arr {
		h.Push(e)
		// Why does this work? Well because the h.Push(e) will heapify, therefore the top element being the k+1 largest is already too small, so we get rid of it
		if h.Len() > k {
			h.Pop()
		}
	}

	v, _ := h.Pop()
	return v
}

// Given an integer array nums and an integer k, return the kth largest element in the array.
// Note that it is the kth largest element in the sorted order, not the kth distinct element.
// Can you solve it without sorting?

// Sorting takes O(nlog(n)) at best, but being able to determine the kth largest items means we need to know all k-1 larger elements by definition.
// so we might not have to look at the whole array. But if we aren't going to sort it, there is no way to know the order without looking through the whole array. So our best time complexity will be O(n).
// But really this is just like looking for the largest element, where we do a comparison every time we look at a new element and compare it with our current largest.
// Except, this time, we will store k variables minimum. And we will also need to determine the order in which to store the variables.

// This is beginning to sound a lot like a heap question. We can instantiate a max heap of size K. Then we go through our array and populate the min heap.
// Iterating through the array takes O(n). Inserting each time into the maxheap will take on average O(log(n)), and taking the top element will take O(1).
// So we can create a minHeap, where the top element will be our kth element, it will be smaller than the k-1 elements underneath it.
// So based on that we have O(n * log(n)) right? No! We don't. Why? Because our heap never gets larger than k, so actually, our running time is O(n * log(k)) and this is actually meaningfully better.
// In the absolute worst case where k = n, then we are no better off than just sorting the array, but in most cases, we are considerably better off.

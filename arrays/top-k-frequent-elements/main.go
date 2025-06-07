package main

import (
	"fmt"

	heap "github.com/idsulik/go-collections/v3/priorityqueue"
)

func main() {
	fmt.Print(TopKFreqEl([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Print(TopKFreqEl([]int{1}, 1))
	fmt.Print(TopKFreqEl([]int{-1, 1, 6, 6, 10, 10, 10}, 3))
}

func TopKFreqEl(arr []int, kk int) []int {
	// We want to make a minHeap, so smaller nodes are higher priority
	h := heap.New(func(n, m *Node) bool {
		return n.Freq < m.Freq
	})

	cnt := map[int]int{}
	for _, e := range arr {
		cnt[e]++
	}

	for k, v := range cnt {
		// Push the current element onto the heap
		h.Push(&Node{
			Val:  k,
			Freq: v,
		})

		// Ensure the heap doesn'g grow larger than k, we only need to ever push one element, since we never let it grow out of control
		if h.Len() > kk {
			h.Pop()
		}
	}

	out := []int{}
	for !h.IsEmpty() {
		v, _ := h.Pop()
		out = append(out, v.Val)
	}
	return out
}

type Node struct {
	Val  int
	Freq int
}

// Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.
// Algorithm must be O(nlog(n))

// Ok if it's O(nlog(n)) that we are aiming for, we should be allowed to sort the array, and we should also be allowed to traverse it once, so let's suppose we have a sorted array that looks something like
// [a,a,a,b,b,c,c,c,c,e,e...,x,x,x,x,x,x,x,x,y,z] in fact, we are told that the array can be up to 10^5 long and contain numbers ranging between -10^4 and 10^4
// We are also told that the answer is unique, meaning we don't have to worry about the case where two numbers so up the same number of time

// So I think a naive solution would be to sort the array first, and make a note of the indexes of the boundaries between numbers, the difference between the start index and end index for a given number is it's frequency
// Then we want to insert this frequency into a data structure that lends itself to easily searching, the nodes of the data structure will be ordered by frequency and they will also have the value of the number
// The idea of using a max heap with a node structure like below does come to mind:

// var Node struct {
// 	Val  int
// 	Freq int
// }

// But I think an equally valid approach is to just store these nodes in an array, and then sort by the frequency which is again an Olog(n) operation, and then to just return the Vals of the first k nodes.
// But we incur a cost of sorting the whole array in this case, when really we only need the first k elements, so is there a way to identify the top k quickly, yes using a heap.
// And I don't even thing there is any point sorting the array at the start, we can just iterate through it once and count each time an element shows up. So actually we only incur O(n) cost to generate a map of elements mapped to frequencies.

// Then we will incur another O(n) to loop over the map and get data into a the heap. Here's the trick, we only care about the top k frequent, so we will make our heap of size k.
// And we will only insert into our heap if the freq of the current element is greater than the smallest element currently in our heap, so we will use a minheap to be able to very quickly pop this element out and push the new one in
// We don't need to handle this logic - minHeap.Top < currNode.Freq then minHeap.Pop then minHeap.Push(currNode) - ourselves since the heap should do it for us

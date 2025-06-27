package main

import (
	heap "github.com/idsulik/go-collections/priorityqueue"
)

func SortLL(ll []LL) LL {
	h := heap.New[*Node](func(a, b *Node) bool {
		return a.Val < b.Val
	})

	for _, l := range ll {
		if l.Head != nil {
			h.Push(l.Head)
		}
	}

	var out LL
	var curr *Node
	for !h.IsEmpty() {
		m, _ := h.Pop()
		if out.Head == nil {
			out.Head = m
			curr = out.Head
		} else {
			curr.Next = m
			curr = m
		}

		if m.Next != nil {
			h.Push(m.Next)
		}
	}

	return out
}

type Node struct {
	Val  int
	Next *Node
}

type LL struct {
	Head *Node
}

// O(Nlog(k)) where n is the number of nodes and k is the size of each LL, because we iterate over all nodes at least once, and we perform a minHeap Pop/Push for all nodes which is log(k)

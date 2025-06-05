package main

import "fmt"

func main() {
	slcs := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}

	for i, slc := range slcs {
		v := LLFromSlc(slc)
		v.Print()
		fmt.Printf("removing element %v from back\n", i+1)
		v.RemoveNthNode(i + 1)
		v.Print()
		fmt.Println()
	}
}

type LL struct {
	Head *LLNode
	Next *LLNode
}

type LLNode struct {
	Val  int
	Next *LLNode
}

func (ll *LL) RemoveNthNode(n int) *LLNode {
	i, j := ll.Head, ll.Head

	// Start first pointer at node n if it exists
	for k := 0; k < n; k++ {
		if j != nil {
			j = j.Next
		} else {
			// We return the head since we would have exited the linkedlist at this iteration, so there is no nth node from the ll to remove when our ll isn't long enough
			// So we return the ll as is with no nodes removed
			return ll.Head
		}
	}

	if j == nil {
		// Since we haven't returned this means we are at the case where the element we want to remove is the head
		ll.Head = ll.Head.Next
		return ll.Head
	}

	for j.Next != nil {
		j = j.Next
		i = i.Next
	}

	i.Next = i.Next.Next
	return ll.Head
}

func LLFromSlc(arr []int) *LL {
	ll := &LL{
		Head: &LLNode{
			Val:  arr[0],
			Next: nil,
		},
		Next: nil,
	}

	prev := ll.Head
	for i := 1; i < len(arr); i++ {
		n := &LLNode{
			Val:  arr[i],
			Next: nil,
		}
		prev.Next = n
		prev = n
	}
	return ll
}

func (ll *LL) Print() {
	n := ll.Head
	for n != nil {
		fmt.Printf("%v ", n.Val)
		n = n.Next
	}
	fmt.Println()
}

// The idea here is simple... if you know the trick. The trick is to have two runners, the first runner runs ahead n steps before the second runner even starts
// Then both runners are incrementented as they run down the linked list, once the first runner hits the end, becuase it was n ahead of the second runner
// the second runner will be in it's rightful place. I'm going to assume that I'm operating on a singly linked list.

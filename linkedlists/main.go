package main

import "fmt"

func main() {
	slcs := [][]int{
		{1, 2, 3, 4, 5},
		{7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		{17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
	}

	// for i, slc := range slcs {
	// 	v := LLFromSlc(slc)
	// 	v.Print()
	// 	fmt.Printf("removing element %v from back\n", i+1)
	// 	v.RemoveNthNode(i + 1)
	// 	v.Print()
	// 	fmt.Println()
	// }

	for _, slc := range slcs {
		v := LLFromSlc(slc)
		fmt.Print("odd even node, initial ll: ")
		v.Print()

		v.OddEvenNode()
		fmt.Print("odd even node, final ll:   ")
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

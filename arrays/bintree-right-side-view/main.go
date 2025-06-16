package main

import (
	deque "github.com/idsulik/go-collections/v3/deque"
)

func main()

func BinTreeRightSideView(root *Node) []int {
	if root == nil {
		return nil
	}

	out := []int{}
	out = append(out, root.Val)

	d := deque.New[*Node](100)
	d.PushBack(root)

	for d.Len() > 0 {
		l := d.Len()

		b, _ := d.PeekBack()
		out = append(out, b.Val)

		for i := 0; i < l; i++ {
			n, _ := d.PopFront()
			if n.Left != nil {
				d.PushBack(n.Left)
			}
			if n.Right != nil {
				d.PushBack(n.Right)
			}
		}
	}

	return out
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func BinTreeFromArr(arr []int) *Node {
	return nil
}

// Level order traversal
// Take rightmost node in each level
// 				1
//		2				4
//	5		6		7
//	  8
//      9

// out == [1], d == [2, 4]
// out == [1, 4], d = [5, 6, 7]
// out == [1, 4, 7]

// Time complexity
// Examine every node once, push it onto deque once, we pop if off deque, we do a constant number of work to check left and right children
// Therefore this seems to be linear in the input size, i.e. the total number of nodes in the tree, tree depth has no effect, so I'll say it's O(n) in time
// For space we initialise an array, the largest the array can ever get is N/2, we want the shortest and widest binary tree possilbe, and this is a balanced tree, where the lowest
// level of the tree has N/2 nodes

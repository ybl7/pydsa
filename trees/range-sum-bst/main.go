package main

func RangeSumBST(n *Node, l, r int) int {
	if n == nil {
		return 0
	}

	if l <= n.Val && n.Val <= r {
		return n.Val + RangeSumBST(n.Left, l, r) + RangeSumBST(n.Right, l, r)
	} else if l > n.Val {
		return RangeSumBST(n.Right, l, r)
	} else {
		return RangeSumBST(n.Left, l, r)
	}
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

package main

func SumRootLeaf(n *Node, num int) int {
	if n == nil {
		return 0
	}
	num = 10*num + n.Val
	if n.Left == nil && n.Right == nil {
		return num
	}
	return SumRootLeaf(n.Left, num) + SumRootLeaf(n.Right, num)
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

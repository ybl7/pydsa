package main

func DiamBinTree(root *Node) int {
	max := 0
	out := DiamBinNode(root, 0, &max)
	return out
}

func DiamBinNode(n *Node, depth int, max *int) int {
	// Recursed onto non existent node, returning input depth
	if n == nil {
		return depth
	}

	// Get left and right subtree depths
	l := DiamBinNode(n.Left, depth+1, max)
	r := DiamBinNode(n.Right, depth+1, max)

	// Update max if required, our l and r are the absolute depths from root to leaf
	// If we just want to find the difference between the two nodes l and r then we need to remove the component
	// that is the depth of the current node from the equation
	curr := (l - depth) + (r - depth)
	if *max < curr {
		*max = curr
	}

	// Take the larger depth of the two subtrees to return to the parent
	if l > r {
		return l
	} else {
		return r
	}
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

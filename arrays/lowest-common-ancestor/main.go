package main

func main()

func LowestCommonAncestor(n *Node, p, q int) *Node {
	if n == nil || n.Val == p || n.Val == q {
		return n
	}

	// If a target node exists in left/right child subtree l/r will be non nil
	l := LowestCommonAncestor(n.Left, p, q)
	r := LowestCommonAncestor(n.Right, p, q)

	// If both are non nil, it means that this node is the LCA of the targets
	// Once we hit this, our subtrees that would have returned two values since both targets are found, now return just the LCA
	// Since we found the LCA, this implies that all other subtrees above the once that is started from LCA will always return nil
	// so we will correctly propagate this value up. This is the crucial step where we replace the node as we discovered it by
	// an LCA for both nodes, if we never find a node that satisfies this, it implies that one target node is the LCA for the other
	if l != nil && r != nil {
		return n
	}

	// Propagate node in the right subtree up, if the left node is found in the left subtree it would have been handled above
	// This necessarily implies that if the other target does exist it will be in the right subtree, or it is in a tree at a higher level
	// The converse applies for the left subtree
	if l == nil && r != nil {
		return r
	}
	if l != nil && r == nil {
		return l
	}
	return nil
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

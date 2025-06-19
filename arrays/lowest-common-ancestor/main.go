package main

func main()

func LowestCommonAncestor(n *Node, p, q int) *Node {
	if n == nil || n.Val == p || n.Val == q {
		return n
	}

	// If we aren't at root, or we didn't find a target value, search the left and right subtrees
	l := LowestCommonAncestor(n.Left, p, q)
	r := LowestCommonAncestor(n.Right, p, q)

	// We found a node matching a value in both subtrees, therefore this node must be the LCA
	// Our function does two things it returns the node when we find either p or q, but then it also returns the LCA
	// Why does this work? Because when you find a node, it IS it's own LCM, but it may not be the LCM for the other target
	// So really this algorithm finds the optimal LCM for a number q, then "de-optimises" it so that it works for p and q
	if l != nil && r != nil {
		return n
	}

	// Left subtree contains both, since right subtree return nil, we exit at the optimal node for either p and q
	// which is necessarily equal to either p or q, which the other value being lower in the tree
	if l != nil && r == nil {
		return l
	}
	// Right subtree contains both since left subtree is empty, take the best node we got from the right subtree
	if l == nil && r != nil {
		return r
	}

	// l == nil && r == nil
	return nil
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// Thought about in a different way, we can think of this as choosing some node n and searching it's left and right subtree
// Then we decide what to do with the results of our search

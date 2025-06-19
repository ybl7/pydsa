package main

func LowestCommonAncestorIII(a, b *Node) *Node {
	p, q := a, b

	for a != b {
		if a.Parent != nil {
			a = a.Parent
		} else {
			a = q
		}
		if b.Parent != nil {
			b = b.Parent
		} else {
			b = p
		}
	}

	return a
}

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

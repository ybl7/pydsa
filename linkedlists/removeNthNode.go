package main

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

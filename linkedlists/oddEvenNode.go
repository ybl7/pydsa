package main

func (ll *LL) OddEvenNode() {
	odd := ll.Head
	sep := ll.Head.Next // sep for separator, the start of the even sublist
	even := ll.Head.Next

	// The odd tail will always point to the start of the even tail and always comes before, so we'll know when we are done when we run out of elements at the end of the even tail
	for even != nil {
		// Point the current odd tail to the new odd tail, which will be the very fist odd node after the end of the even sublist
		odd.Next = even.Next
		// Set odd to the new tail, if one exists
		if odd.Next != nil {
			odd = odd.Next
		} else {
			// There is no next node so we have reached the end of the ll, so just point odd to the start of the even sublist and return
			odd.Next = sep
			return
		}

		// Point current even tail to new even tail, if it exists it will be at odd.Next, but if odd itself is nil, we have reached the end
		even.Next = odd.Next
		// Set even tail
		even = even.Next

		// Set the odd tail's next to the even tail head
		odd.Next = sep
	}
}

// This question is doable because we are guaranteed a pattern odd, even, odd, even - or rather, we assume it
// Therefore if we visit an even node, we know the next one will be odd and vice versa
// Therefore we will never have an accumilation of say odd, odd, odd, even, this would be difficult to deal with since it's very hard to move
// mulitple nodes in bulk in a linked list, but one node at a time, we can do that.
// Let's instantiate two pointers, i and j, to mark the heads of the odd and even lists respectively

// i == 0, j == 1
// 1 -> 2 -> 3 -> 4 -> 5 -> 6
// Next we will perform the following operations to group the odd and even nodes
// 1.next = 3, 2.next = 4, 3.next = 2
//  - - - - -
// |         |
// |         v
// 1    2 <- 3    4 -> 5 -> 6
//      |         ^
//      |         |
//       - - - - -
// When we write this in a simpler way it's the same as:
// 1 -> 3 -> 2 -> 4 -> 5 -> 6

// i == 0, j == 2
// Let's do one more iteration to really get the hang of it
// 3.next = 5, 4.next = 6, 5.next = 2
//       - - - - - - -
//      |              |
//      |              v
// 1 -> 3    2 -> 4 <- 5    6
//           ^    |    |    ^
//           |    |    |    |
//           |     - - + - -
//             - - - -
// So what we notice is that:
// We always point the tail of the current odd/even node to the next odd/even node
// We always point the tail of the odd node to the start of the even sublist
// The start of the even sublist never changes, the tail of the odd sublist does
// So I think it makes sense to keep pointers to the tails of both sublists, and the head of the even sublist

// We'll keep a pointer n to the next unprocessed node, when n == nil we knoe we are done as we have processed all nodes

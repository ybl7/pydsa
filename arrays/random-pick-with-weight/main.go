package main

import "math/rand"

func RandWeightPick(weights []int) int {
	out := []int{}
	sum := 0
	for _, e := range weights {
		sum += e
		out = append(out, sum)
	}

	rng := rand.Intn(sum) + 1
	l, r := 0, len(out)-1
	for l < r {
		m := (l + r) / 2

		if out[m] < rng {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

// Suppose we have some array [1, 2, 3], we can represent this in anothe array [0, 1, 1, 2, 2, 2]
// If we select randomly from this array there is a 1/6 chance to get index 0, 2/6 to get index 1, and 3/6 to get index 2
// What happens if I have an array like [1, 2, 300000000000000000000000000000000000000000]? If we try to represent this as an array, we run out of memory.
// So how can we represent relative probabilities without enumerating all of the outcomes?
// So actually if we pick an index randomnly from the interval [0, 5] it could fall in one of the following ranges [0, 0], or [1, 2], or [3, 5].
// And as it so happens, the sizes of these ranges is proportional to the probability of them being picked.
// From this image it's clear that the size of the interval is proportional to the probability | 0 | 1 1 | 2 2 2 |
// So a neat trick to be able to represent the size of the ranges, without having to create a suitable array with sufficient elements, for this an array like [1, 3, 6] works.
// Which is interpreted according to the following rule, a number belongs to the smallest element that is greater than it.
// Suppose we pick a random number from 1 to 6, if the number is 4, then it belongs to index 2 since 6 > 4 and is the tightest upper bound.
// This way we can maintain the rule that a range's size is proportional to it's probability without having to specify it explicitly.
// But it's not just proportional, the range also maps back to the index. The final piece of the puzzle is how then we should go about searching through our range arr e.g. [1, 3, 6]
// and we'll note that it is monotonically increasing, which is like saying it is ordered, so we can use binary search to very quickly place a random number in a range and return it's index.

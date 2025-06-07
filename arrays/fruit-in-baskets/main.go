package main

import "fmt"

func main() {
	fmt.Println(FruitInBasket([]int{1, 2, 1}))                         // 3
	fmt.Println(FruitInBasket([]int{0, 1, 2, 2}))                      // 3
	fmt.Println(FruitInBasket([]int{1, 2, 3, 2, 2}))                   // 4
	fmt.Println(FruitInBasket([]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4})) // 5
	fmt.Println(FruitInBasket([]int{1, 1, 2, 2, 3, 3, 2, 2}))          // 6
	fmt.Println(FruitInBasket([]int{0, 1, 6, 6, 4, 4, 6}))             // 5
}

func FruitInBasket(arr []int) int {
	if len(arr) <= 2 {
		return len(arr)
	}

	// Key is fruit type, value is last seen index
	m := make(map[int]int)
	var max int
	var curr int

	for i, j := 0, 0; j < len(arr); j++ {
		// Add everything to the map, we'll ensure that we only keep two keys later
		m[arr[j]] = j

		// Remove the minimum element when we have 3 elements
		if len(m) > 2 {
			minK := MinMapVal(m)
			i = m[minK] + 1
			delete(m, minK)
		}

		curr = j - i + 1
		if max < curr {
			max = curr
		}
	}
	return max
}

func MinMapVal(m map[int]int) int {
	var min int
	first := true
	for k, v := range m {
		if first {
			min = k
			first = false
		} else {
			if v < m[min] {
				min = k
			}
		}
	}
	return min
}

// This question is very wordy but it essentially boils down to finding a longest contiguous subarray with only integers
// We are just given the array and we don't know the elements that will be in the array ahead of time and we are told we can have up to 10^5 elements
// Let's take an illustrative example [a,b,c,b,a,d,a,c,d,e,b,a,s,t,b,a,b,a,b,c,d,e]
// This is a question of maxima, so we don't need to remeber ALL possible subsets, we just need to be able to remember the best subset
// This means we can discard old information once it becomes irrelevant

// So we can take a sliding window approach, where we increase the window to the right until we hit a fruit that is not of the current two types in the window
// When we do this we push the start of the window forward until we only have two fruits in it again
// Let's look at an example to understand this, the window is delimited by the | symbols

// Iteration 1
// [a,b,c,b,a,d,a,c,c,d,e,b,a,s,t,b,a,b,a,b,c,d,e]
//  | |

// Iteration 2
// [a,b,c,b,a,d,a,c,c,d,e,b,a,s,t,b,a,b,a,b,c,d,e]
//    | |

// Iteration 3
// [a,b,c,b,a,d,a,c,c,d,e,b,a,s,t,b,a,b,a,b,c,d,e]
//    |   |

// And so on, so it's fairly easy to get the end of the window, we just keep moving it forward by one at a time, and so long as we don't have > 2 types of value we continue doing this
// But what about updating the start of the window, how do we know where to move it to?

// Suppose at some point we get a valid window like (assume it starts from 0 i.e. the index of the first c is )
// c b c c b c b c c c
// |                 |
// Then we encounter some new tree with fruit a
// Up until we hit the a, our window would have looked something like
// c b c c b c b c c c a
// |                 |
// Now it's clear to see that our new window should be, we want to get the index of the first c in our new window
// c b c c b c b c c c a
//				 |     |
// The question is, how do we get the index of the first valid c? Our previous window had a start that pointed to the first c and and end that pointed to the last c
// Well it's actually just the index of the last valid b (which is no invalid!) added to one

// Let's take it back a step before we added the a. Let's keep two invariants, the last valid indexes for the current fruits in the window, and call them lc and lb respectively, we'd end up with
// c b c c b c b c c c
// |                 |
// lc = 9
// lb = 6

// Then as soon as we introduce the a, we check which fruit we need to eliminate, in this case b. This takes a moment to understand, but by tracking the last index of every fruit. We know which friut cannot be valid.
// Because if we have two fruits b and c, if the last index of b < c then it must have come before and never showed up again before we encouter a, this necessarily implies that been the last valid index of b and a, we only have c's.
// But by tracking the last valid b, we know EXACTLY where the b's end, at index lb, therefore we know we just need to move our start of sliding window pointer to lb + 1, which by definition must be the start of the c's.
// c b c c b c b c c c a
//               |     |
// lc = 9
// la = 10

// So our strategy will be to
// (a) keep track of the current fruits in the window and their last valid index
// (b) increment the back of the window by one, do this without any checks until we have 3 elements in the store
// (c) when we have 3 elements, mapped back to their indexes, we just remove the one with earliest index and bump i to that index+1 and calculate the window size which will be j - i + 1 since we want the inclusive length
// This question took way too long :/

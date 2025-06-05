package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	numsArrs := [][]int{
		{-1, 0, 1, 2, -1, -4},
		{1, 2, 3, 4, 5},
		{0, 0, 0, 0},
		{-4, -1, -1, 0, 1, 2, 2},
		{-10, -7, -3, -1, 0, 3, 7, 10},
		{-3, -5, -7, -9},
	}

	for i, nums := range numsArrs {
		fmt.Printf("%d.\tnums: [", i+1)
		for j, num := range nums {
			fmt.Printf("%d", num)
			if j < len(nums)-1 {
				fmt.Print(", ")
			}
		}
		fmt.Println("]")

		triplets := ThreeSum(nums)
		fmt.Print("\n\tTriplets: [")
		for j, triplet := range triplets {
			fmt.Printf("[%d, %d, %d]", triplet[0], triplet[1], triplet[2])
			if j < len(triplets)-1 {
				fmt.Print(", ")
			}
		}
		fmt.Println("]")
		fmt.Println(strings.Repeat("-", 100))
	}
}

func ThreeSum(arr []int) [][]int {
	// Sort first since our algorithm relies on it
	sort.Ints(arr)
	out := [][]int{}

	for k := 0; k < len(arr)-2; k++ {
		// If the current number is greater than 0, break the loop, the array is sorted, so if we start positive we will never go to zero since i and j both larger than k
		if arr[k] > 0 {
			break
		}

		// The current number is either the first element or not a duplicate of the previous element
		if k == 0 || arr[k] != arr[k-1] {
			// i is just right of current element, j is always the maximal value
			i := k + 1
			j := len(arr) - 1

			for i < j {
				sum := arr[k] + arr[i] + arr[j]

				if sum == 0 {
					o := []int{arr[k], arr[i], arr[j]}
					out = append(out, o)

					// increment i if there are duplicates, the i < j serves as a short circuit to prevent us from overincrementing i
					for i < j && arr[i+1] == arr[i] {
						i++
					}
					// decrement j if there are duplicates, the i < j serves as a short circuit to prevent us from overdecrementing j
					for i < j && arr[j-1] == arr[j] {
						j--
					}
					// final increment/decrement to get onto a non duplicate
					i++
					j--
					// We don't have to worry about duplication here since we handle it if there is a viable triplet, if not we just iterate more than we need to over dupliate i and j, but it doesn't ever show up in the final slice
				} else if sum < 0 {
					i++
				} else if sum > 0 {
					j--
				}
			}

		}
	}
	return out
}

// With 3 sum we could take a similar strategy to the one we took for 2 sum
// That is we take our target which would now be t = a + b + c, we then set a new target t - a = t' = b + c
// Then we treat t' as a target for a 2 sum problem as we did before, recall that 2 sum using a single hashmap ran in O(n) time and O(n) space
// In our case we will need to run 2 sum n times, we can save a little time and space by only constructing the map once and marking the current elemnt a so that it isn't counted
// So the solution might look like looping over the array one time and constructing a map containing all elements mapping to their index
// Then for each element of the map, we take that our 3 sum target less that element and use this as the target for 2 sum over all other elements, we'll do this n times
// So we have O(n) time and O(n) space to construct the initial map, but then O(n) time and space for each element e, so our overall complexity will be O(n^2)

// We can do better, ordering will help us a lot here, since if elements are ordered, it's very easy to see when your sum is too small or too great
// Suppose we order our initial array, then we take an element as our a, that is the first element in our sum a + b + c = t, the complement we are looking for is t - a
// We will hold a constant, then explore the rest of the array for a viable pair to go with a.
// The trick is to intelligently search through the combinations while discarding those that are obviously not vialble. Intelligently you say, but how?
// If we set a pointer to be the first element to the right of a, then this must be the next smallest element.
// Conversely if we set a pointer to be the rightmost value, this will be the largest value.

// The cool part is that by having i and j at the extremes, we can then move the pointers inwards without ever missing a potential suitable pair.
// j is necessarily the max, moving j down will always make the sum smaller, the converse applies with i.
// So we will take our sum a + arr[i] + arr[j] and get an answer, it may be equal to, smaller than, or greater than t
// The question itself puts a more stringent constraint that the sum must be 0 and not just any old target, let's see why this helps us.

// Suppose my array is something like [-4, -1, -1, 0, 2, 2] and let's initially hold k = -4
// Let's draw a table of all possible combinations and their sums to get a better understanding, i on the x axis and j on the y axis.
// We place the value of the sum of i and j where possible and leave blank where it's not since i and j can't be the same index simultaneously

//       -1 -1  0  2  2
//
// -1       -2 -1  1  1
// -1    -2    -1  1  1
//  0  	 -1 -1     2  2
//  2  	  1  1  2     4
//  2     1  1  2  4

// And because
// we don't care about which of the indexes i or j produce the sum,
// and we also don't care about indexes (unlike two sum), rather just unique values that sum to 0
// and j will always be greater than i, we can simplify the table quite a bit

//     	 -1 -1  0  2  2
//
// -1
// -1  	 -2
//  0  	 -1 -1
//  2  	  1  1  2
//  2  	  1  1  2  4

// So this is our search space once we place the contraint that i < j, i and j can't be the same index at any time
// Now if we initialise i and j as described, this puts us in the bottom left corner of the triangle with a value 1
// Let's consider the following cases

// sum == 0
// We have found a pair k, i, j so we store the values, there may be other pairs however that we must look for
// We increment i until we reach a new value (avoiding duplicates), and decrement j until we reach a new value
// This is because if we change i without changing j, or vice versa, we will necessarily be too large/small repsecitively, since we are perfectly 0 with the current values
// So our hope is to move j down and i up so that the increase in i is sufficiently offset by the decrease in j

// sum < 0
// Since we have undershot, if we were to decrement j while keeping i the same, we would only get smaller sums - this value of i will never produce a large enough sum with smaller j.
// Since we start the problem with the maximal value of j, the current i value will be too small with ANY other value of j since j can only get smaller.
// Therefore we can safely discard this value of i, knowing it is not large enough to ever produce a sum == 0. But the next value of i may.

// sum > 0
// Similarly when the sum is too large for a given j, this means that if we increment i, we will only ever get bigger sums.
// Since we start the problem with the minimal value of i, and i can only get larger, we can only grow the sum from. So it's safe to discard this value of j.

// Now the final thing to say is that once we move i and j, these indexes become the new smallest/largest viable indexes, because we ruled out smaller/larger values respectively before moving
// That was for a fixed k, if we initialise k to be the first element, then after our first search, we have necessarily exhausted all options containing k = 0
// Therefore we can increment k without having to look back at it for the next value of k, if a combination with the previous and current k exists, we could have already found it
// So we save ourselves some work, this is like saying if I discovered that 3 * 5 = 15 I don't need to check that 5 * 3 = 15 when it's 5's turn

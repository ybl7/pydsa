package main

func SubArrSumK(arr []int, tgt int) int {
	cnt := 0
	sum := 0
	// key is the contiguous sums, value is the number of occourences, initialise to 0 for case where first index is
	subArrSums := map[int]int{
		0: 1,
	}

	for _, e := range arr {
		sum += e
		subArrSums[sum]++

		if c, ok := subArrSums[sum-tgt]; ok {
			cnt += c
		}
	}

	return cnt
}

// You might think a sliding window approach could work, but this breaks because the numbers can be negative
// So you have no idea if growing/shrinking the window will make the sum bigger or smaller, so there REALLY is
// No way around considering EVERY contiguous subset, becuase suppose I shrink the window, perhaps there was a number just
// So the best we can do here is not re-calculate the sums all over again

// This is more a maths problem than a DSA problem, given some array arr = [a,b,c,d,e,f,g]
// Take the current index as d, there are 4 continguous subarrs that end at d
// [a,b,c,d], [b,c,d], [c,d], [d]
// Sum([a,b,c,d]) = Sum(a) + Sum(b) + Sum(c) + Sum(d)
// How can we get these sums without calculating each time?

// Well notice how if we just traverse the array and note the sum of the subarray up until the cuurent index we get:
// [a, a+b, a+b+c, a+b+c+d], ok but how does this help up us get the subarrs that end at d?
// Well becuase a+b+c+d - a is b+c+d, so we only need to track the sum of the sub arr at each index and we can
// always calculate the sum of subarray ending at d

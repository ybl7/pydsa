package main

func MaxFreqAfterSubArrOp(arr []int, k int) int {
	kCnt := 0
	for _, e := range arr {
		if e == k {
			kCnt++
		}
	}

	cnt := 0
	// We loop from 1->50 since we are told the numbers in the array have this range
	// so there are 50 possibilites that each index can takae
	for x := 1; x < 51; x++ {
		pos := Kadane(arr, k, x)  // el + x = k
		neg := Kadane(arr, k, -x) // el - x = k
		if pos > cnt {
			cnt = pos
		}
		if neg > cnt {
			cnt = neg
		}
	}

	return kCnt + cnt
}

func Kadane(arr []int, k, x int) int {
	t := k - x
	cnt := 0
	max := 0

	for _, e := range arr {
		if e == t {
			cnt++
		} else if e == k {
			cnt--
		}
		if cnt < 0 {
			cnt = 0
		}
		if cnt > max {
			max = cnt
		}
	}

	return max
}

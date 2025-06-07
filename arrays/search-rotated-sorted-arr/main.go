package main

import "fmt"

func main() {
	fmt.Println(SchRotSortArr([]int{4, 5, 6, 7, 0, 1, 2}, 0, 0))       // 4
	fmt.Println(SchRotSortArr([]int{4, 5, 6, 7, 0, 1, 2}, 0, 3))       // -1
	fmt.Println(SchRotSortArr([]int{1}, 0, 0))                         // -1
	fmt.Println(SchRotSortArr([]int{4, 5, 6, 7, 8, 9, 1, 2, 3}, 0, 2)) // 7
}

func SchRotSortArr(arr []int, offset, tgt int) int {
	if len(arr) == 0 {
		return -1
	}
	if len(arr) == 1 {
		if arr[0] == tgt {
			return offset
		}
	}

	m := len(arr) / 2
	end := len(arr) - 1

	if arr[0] > arr[end] {
		lArr := arr[:m]
		rArr := arr[m:]

		// Search both sub arrays
		lIdx := SchRotSortArr(lArr, offset, tgt)
		rIdx := SchRotSortArr(rArr, offset+m, tgt)

		if lIdx != -1 {
			return lIdx
		}
		if rIdx != -1 {
			return rIdx
		}
		return -1
	} else {
		b := BinSch(arr, tgt)
		if b != -1 {
			return offset + b
		}
		return -1
	}
}

func BinSch(arr []int, tgt int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == tgt {
			return mid
		}

		if arr[mid] < tgt {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

// Obvious solution is O(n) just iterate the array, but the question requires us to do it in O(logn) or better
// It's sorted, so it lends itself to binary search, but it's also rotated so it doesn't
// But we can still use binary search, suppose we have some array [1,2,3,4,5,6,7,8,9] and we rotate it [4,5,6,7,8,9,1,2,3]
// If we could find the point of rotation we could apply binary search i.e we could do BinSearch([4,5,6,7,8,9]) and BinSearch([1,2,3])
// The problem is if loop through the array to try and find this point, we are already O(n)

// But here's the trick, suppose we just randomly split the array in two: [4,5,6,7,8] amd [9,1,2,3]
// We know that the array is ordered, just rotated, so if we check the two endpoints of the arrays we will immediately know if it could contain our element
// If we check the endpoints of the subarray we will also immediately know if this array even contains our number
// If arr[end] > arr[start], then we know this is an ordered array, but note that we can discard the array it if arr[start] > tgt ot arr[end] < tgt
// if the array actually contains our number (after the check against endpoints), we can use standard binary search on it
// If arr[end] < arr[start], then we just split the array in half again. This time to get [9,1], [2,3].

// So one subroutine splits arrays if arr[end] < arr[start], this happends recursively until we hit the base case
// In the worst case scenario the array will be rotated by just 1, so we will have to run this subroutine log(n) times doing constant work each time to check the endpoints before the next split
// As for the ordered part of the subarray, the running time is just standard binary search whihc is O(log(n))
// So in any given recursive call, you will at most call both (a) call the split subroutine, and (b) call binary search
// The split subroutine always does constant work at every level of the redursive tree since it all it does is check the enpoints of the array,
// the size of the array has no bearing on it, so we just sum a constant log(n) times, which is still just log(n) work
// So (before writing any code) it looks like our running time will be the sum of log(n) + log(n) which is still just O(log(n))

// I'm going to assume I don't run into the case where the endpoints are the same value, in fact the question even says that the values in the array are distinct

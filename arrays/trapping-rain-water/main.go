package main

import (
	"fmt"
	"math"
)

func TrapRain(arr []int) int {
	mL := arr[0]
	mR := arr[len(arr)-1]
	maxL := []int{mL}
	maxR := []int{mR}

	for i := 1; i < len(arr); i++ {
		if arr[i] > mL {
			mL = arr[i]
		}
		maxL = append(maxL, mL)
	}

	for j := len(arr) - 2; j >= 0; j-- {
		if arr[j] > mR {
			mR = arr[j]
		}
		maxR = append(maxR, mR)
	}
	// Reverse maxL since we appended instead of prepening
	for i, j := 0, len(maxR)-1; i < j; {
		maxR[i], maxR[j] = maxR[j], maxR[i]
		i++
		j--
	}

	sum := 0
	for i, c := range arr {
		s := math.Min(float64(maxR[i]), float64(maxL[i])) - float64(c)
		if s > 0 {
			sum += int(s)
		}
	}
	return sum
}

func main() {
	fmt.Println(TrapRain([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

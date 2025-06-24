package main

func Pow(x float64, n int) float64 {
	// Inner function to perform the quick exponentiation.
	quickPower := func(base float64, exp int) float64 {
		result := 1.0
		for exp != 0 {
			// Odd exponent, store base^1 in result, we will lose it later in exp >>= 1
			// Only encountered twice, at the start if the exponent is odd, at the end when we reach exponent 1
			if exp&1 != 0 {
				result *= base
			}
			// Exp is at least two, so we can square the base then halve it
			base *= base
			exp >>= 1
		}
		return result
	}

	if n >= 0 {
		return quickPower(x, n)
	} else {
		return 1 / quickPower(x, -n)
	}
}

// This is better than an O(n) approach where you multiply the base in each iteration by itself
// You approach the exponent logarithmically since you are halving the exponent each iteration, so this is O(log(n))

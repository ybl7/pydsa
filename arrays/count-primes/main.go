package main

func CountPrimes(n int) int {
	if n < 2 {
		return 0
	}

	primes := make([]bool, n)
	for i := range primes {
		primes[i] = true
	}

	out := 0
	for i := 2; i < n; i++ {
		if primes[i] {
			out++
			// Why start from i * i == i ^ 2, because all numbers between i and i^2 will be i * something_smaller_than_i, and the prior step would have already marked them as non prime
			for j := i * i; j < n; j += i {
				primes[j] = false
			}
		}
	}

	return out
}

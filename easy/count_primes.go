package easy

func CountPrimes(n int) int {
	if n <= 2 {
		return 0
	}

	primes := []int{2}
	for i := 3; i < n; i += 2 {
		if IsPrime(i, primes) {
			primes = append(primes, i)
		}
	}

	return len(primes)
}

func IsPrime(n int, primes []int) bool {
	for _, x := range primes {
		if x*x > n {
			break
		}
		if n%x == 0 {
			return false
		}
	}
	return true
}

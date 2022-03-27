// https://leetcode-cn.com/problems/count-primes/
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

func CountPrimes2(n int) int {
	if n <= 2 {
		return 0
	}
	isComposite := make([]bool, n)
	count := n / 2                   // 去掉所有偶数，不要1要2
	for i := 3; i*i < n; i = i + 2 { // 偶数一定不是质数，且只用考虑小于开方数的质因子
		if !isComposite[i] {
			for j := i * i; j < n; j = j + 2*i { // 避免偶数和重复遍历
				if !isComposite[j] {
					isComposite[j] = true
					count--
				}
			}
		}
	}
	return count
}

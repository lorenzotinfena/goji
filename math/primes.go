package math

/*
Time complexity: O(n)

Optimized version of the following:

	func genPrimesBest(n int) []bool {
		isPrime := make([]bool, n+1)
		for i := 2; i <= n; i++ {
			isPrime[i] = true
		}
		primes := make([]int, 0)
		var f func(int, int)
		f = func(i int, prod int) {
			for j := 0; j <= i && primes[j]*prod <= n; j++ {
				isPrime[primes[j]*prod] = false
				f(j, primes[j]*prod)
			}
		}
		lastIndex := 0
		for i := 2; i < n; i++ {
			if isPrime[i] {
				primes = append(primes, i)
				f(lastIndex, i)
				lastIndex++
			}
		}
		return isPrime
	}
*/
func PrimesUpTo(n int) []bool {
	isPrime := make([]bool, n+1)
	isPrime[2] = true
	for i := 3; i <= n; i += 2 {
		isPrime[i] = true
	}
	primes := []int{2}
	var f func(int, int)
	f = func(i int, prod int) {
		for j := 0; j <= i && primes[j]*prod <= n; j++ {
			isPrime[primes[j]*prod] = false
			f(j, primes[j]*prod)
		}
	}
	lastIndex := 1
	for i := 3; i < n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
			f(lastIndex, i)
			lastIndex++
		}
	}
	return isPrime
}

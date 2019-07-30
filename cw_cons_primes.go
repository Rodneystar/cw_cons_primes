package cw_cons_primes

func genPrimesBelow(n int) []int {
	primes := make([]int, 0)
	primes = append(primes, 0, 1, 2, 3)
	for i := 5; i < n; i += 2 {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func isPrime(n int) bool {
	for i := 2; i < n/2+1; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

var primes []int

func ConsecKprimes(k int, arr []int) int {
	primes = genPrimesBelow(10000)
	// your code
	return 0
}

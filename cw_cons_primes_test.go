package cw_cons_primes

import (
	"testing"
)

func dotest(t *testing.T, k int, arr []int, exp int) {
	var ans = ConsecKprimes(k, arr)
	if ans != exp {
		t.Errorf("%d should equal %d", ans, exp)
	}
}

func Test_First(t *testing.T) {
	dotest(t, 2, []int{10081, 10071, 10077, 10065, 10060, 10070, 10086, 10083, 10078, 10076, 10089, 10085, 10063, 10074, 10068, 10073, 10072, 10075}, 2)
	dotest(t, 6, []int{10064}, 0)
	dotest(t, 1, []int{10054, 10039, 10053, 10051, 10047, 10043, 10037, 10034}, 0)
	dotest(t, 3, []int{10158, 10182, 10184, 10172, 10179, 10168, 10156, 10165, 10155, 10161, 10178, 10170}, 5)
	dotest(t, 2, []int{10110, 10102, 10097, 10113, 10123, 10109, 10118, 10119, 10117, 10115, 10101, 10121, 10122}, 7)
}

func Test_isPrime(t *testing.T) {
	t.Log(isPrime(5))
}

func Test_genPrimes(t *testing.T) {
	t.Log(genPrimesBelow(10000))
}

package int

import (
	"strconv"
)

func max(nums ...int) int {
	m := nums[0]
	for _, i := range nums {
		if m < i {
			m = i
		}
	}
	return m
}

func min(nums ...int) int {
	m := nums[0]
	for _, i := range nums {
		if m > i {
			m = i
		}
	}
	return m
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func pow(x, y int) int {
	res := 1
	for i := 0; i < y; i++ {
		res *= x
	}
	return res
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func toNums(str string) []int {
	var nums []int
	for _, r := range str {
		n, _ := strconv.Atoi(string(r))
		nums = append(nums, n)
	}
	return nums
}

func divisor(n int) []int {
	res := []int{}
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			res = append(res, i)
			if i*i != n {
				res = append(res, n/i)
			}
		}
	}
	return res
}

func primeTable(n int) map[int]bool {
	prime := make(map[int]bool)
	for i := 0; i <= n; i++ {
		prime[i] = true
	}
	prime[0] = false
	prime[1] = false
	for i := 2; i*i <= n; i++ {
		if !prime[i] {
			continue
		}
		for j := i + i; j <= n; j += i {
			prime[j] = false
		}
	}
	return prime
}

func primeFactors(n int) map[int]int {
	pfs := make(map[int]int)
	for n%2 == 0 {
		pfs[2]++
		n = n / 2
	}
	for i := 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			pfs[i]++
			n = n / i
		}
	}
	if n > 2 {
		pfs[n]++
	}

	return pfs
}

func modPow(a, n, mod int) int {
	res := 1
	for n > 0 {
		if (n & 1) > 0 {
			res = res * a % mod
		}
		a = a * a % mod
		n >>= 1
	}
	return res
}

func modInv(a, mod int) int {
	b := mod
	u := 1
	v := 0
	for b > 0 {
		t := a / b
		a -= t * b
		a, b = b, a
		u -= t * v
		u, v = v, u
	}
	u %= mod
	if u < 0 {
		u += mod
	}
	return u
}

// Combination
const MAX = int(1e+6)
const MOD = int(1e+9) + 7

var fac [MAX]int
var finv [MAX]int
var inv [MAX]int

func initComb() {
	fac[0], fac[1] = 1, 1
	finv[0], finv[1] = 1, 1
	inv[1] = 1
	for i := 2; i < MAX; i++ {
		fac[i] = fac[i-1] * i % MOD
		inv[i] = MOD - inv[MOD%i]*(MOD/i)%MOD
		finv[i] = finv[i-1] * inv[i] % MOD
	}
}

func modComb(n, k int) int {
	if n < k {
		return 0
	}
	if n < 0 || k < 0 {
		return 0
	}
	return fac[n] * (finv[k] * finv[n-k] % MOD) % MOD
}

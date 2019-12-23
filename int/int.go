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

func modComb(big, small, mod int) int {
	if big < 0 || small < 0 || big < small {
		return 0
	}
	tmp := modFact(big, mod) * modInv(modFact(small, mod), mod) % mod
	return tmp * modInv(modFact(big-small, mod), mod) % mod
}

func modFact(a, mod int) int {
	res := 1
	for i := a; i > 0; i-- {
		res *= i
		res %= mod
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

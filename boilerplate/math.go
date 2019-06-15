package main

// n factorial (mod M)
// return: array of factorial of 0-n
func mfactorial(n, M uint64) []uint64 {
	fact := make([]uint64, n+1)
	fact[0] = 1
	for i := uint64(1); i <= n; i++ {
		fact[i] = (fact[i-1] * i) % M
	}
	return fact
}

// x^n (mod M)
// See https://qiita.com/ofutonfuton/items/92b1a6f4a7775f00b6ae
func mpow(x, n, M uint64) uint64 {
	ans := uint64(1)
	for n != 0 {
		if n&1 == 1 {
			ans = ans * x % M
		}
		x = x * x % M
		n = n >> 1
	}
	return ans
}

// nCk (mod M)
// @see https://qiita.com/ofutonfuton/items/92b1a6f4a7775f00b6ae
func mcombi(n, k, M uint64) uint64 {
	if n == 0 && k == 0 {
		return 1
	} else if n < k || n < 0 {
		return 0
	}
	fact := mfactorial(n, M)

	ans := fact[n] * mpow(fact[k], M-2, M) % M
	ans *= mpow(fact[n-k], M-2, M) % M

	return ans % M
}

func pow(x, y int) int {
	//_, _ = fmt.Printf("pow: %v, %v\n", x,y)
	z := x
	if y == 0 {
		return 1
	}
	for i := 1; i < y; i++ {
		z = z * x
	}
	return z
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

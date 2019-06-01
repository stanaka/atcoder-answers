package main

import (
	"fmt"
	"io"
	"os"
)


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
func combi(n, k, M uint64) uint64 {
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

func answer(reader io.Reader, writer io.Writer) {
	var n,m,k uint64
	_, _ = fmt.Fscan(reader, &n, &m, &k)

	mod := uint64(1000000007)
	var sumX, sumY uint64
	c := uint64(1)
	if k > 2 {
		//c = combination(n*m-2, k-2, mod)
		c = combi(n*m-2, k-2, mod)
		c %= mod
	}
	for i := uint64(1); i < n; i++ {
		sumX += i * (n - i)
		sumX %= mod
	}
	sumX = sumX * m * m
	sumX %= mod
	for i := uint64(1); i < m; i++ {
		sumY += i * (m - i)
		sumY %= mod
	}
	sumY = sumY * n * n
	sumY %= mod
	sum := (sumX + sumY) * c
	sum %= mod
	_, _ = fmt.Fprintf(writer, "%d", sum)
}

func main() {
	answer(os.Stdin, os.Stdout)
}

package main

import (
	"bufio"
	"fmt"
	"io"
)

func primeFactors(n uint64) (pfs map[uint64]uint64) {
	pfs = make(map[uint64]uint64)
	for n%2 == 0 {
		pfs[2]++
		n = n / 2
	}

	var i uint64
	for i = 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			pfs[i]++
			n = n / i
		}
	}

	if n > 2 {
		pfs[n]++
	}
	return
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func GCD(a, b uint64) uint64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b uint64, integers ...uint64) uint64 {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

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

type combi struct {
	fac  []int
	finv []int
	inv  []int
	mod  int
}

func (c *combi) init(n int, m int) {
	c.fac = make([]int, n)
	c.finv = make([]int, n)
	c.inv = make([]int, n)
	c.mod = m

	c.fac[0] = 1
	c.fac[1] = 1
	c.finv[0] = 1
	c.finv[1] = 1
	c.inv[1] = 1
	for i := 2; i < n; i++ {
		c.fac[i] = c.fac[i-1] * i % c.mod
		c.inv[i] = c.mod - c.inv[c.mod%i]*(c.mod/i)%c.mod
		c.finv[i] = c.finv[i-1] * c.inv[i] % c.mod
	}
}

func (c *combi) combi(n, k int) int {
	if n < k {
		return 0
	}
	if n < 0 || k < 0 {
		return 0
	}
	return c.fac[n] * (c.finv[k] * c.finv[n-k] % c.mod) % c.mod
}

func answer(reader io.Reader, writer io.Writer) {
	var c combi
	mod := int(1e9) + 7
	c.init(n+1, mod)
	b := c.combi(k-1, i)
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

func min(a ...int) int {
	res := a[0]
	for i := 1; i < len(a); i++ {
		if res > a[i] {
			res = a[i]
		}
	}
	return res
}

func max(a ...int) int {
	res := a[0]
	for i := 1; i < len(a); i++ {
		if res < a[i] {
			res = a[i]
		}
	}
	return res
}

func minUint64(a ...uint64) uint64 {
	res := a[0]
	for i := 1; i < len(a); i++ {
		if res > a[i] {
			res = a[i]
		}
	}
	return res
}

/*
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
*/

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// uint8  : 0 to 255
// uint16 : 0 to 65535
// uint32 : 0 to 4294967295
// uint64 : 0 to 18446744073709551615
// int8   : -128 to 127
// int16  : -32768 to 32767
// int32  : -2147483648 to 2147483647
// int64  : -9223372036854775808 to 9223372036854775807

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

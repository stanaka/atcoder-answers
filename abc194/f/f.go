package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

const bufferSize = 1024

func nextInt() int {
	if !sc.Scan() {
		panic(nil)
	}
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func nextString() string {
	if !sc.Scan() {
		panic(nil)
	}
	return sc.Text()
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

var production bool

func debugf(format string, v ...interface{}) {
	if !production {
		fmt.Printf(format, v...)
	}
}

func answer(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	buf := make([]byte, bufferSize)
	sc.Buffer(buf, 1e+7)
	sc.Split(bufio.ScanWords)
	s := nextString()
	l := len(s)
	n := make([]int, l)
	for i, x := range s {
		d, _ := strconv.ParseInt(string([]rune{x}), 16, 0)
		n[i] = int(d)
	}
	k := nextInt()
	m := int(1e9 + 7)

	c := make(map[int]struct{})
	dp := make([][]int, l+1)
	for i := range dp {
		dp[i] = make([]int, 17)
	}
	dp[1][1] = n[0] - 1
	for i := 1; i < l; i++ {
		for j := 1; j <= 16; j++ {
			dp[i+1][j] += dp[i][j] * j
			dp[i+1][j] %= m
			if j < 16 {
				dp[i+1][j+1] += dp[i][j] * (16 - j)
				dp[i+1][j+1] %= m
			}
		}
		dp[i+1][1] += 15
		dp[i+1][1] %= m
		c[n[i-1]] = struct{}{}
		lc := len(c)
		for j := 0; j < n[i]; j++ {
			if _, ok := c[j]; ok {
				dp[i+1][lc]++
				dp[i+1][lc] %= m
			} else {
				dp[i+1][lc+1]++
				dp[i+1][lc+1] %= m
			}
		}
	}

	c[n[l-1]] = struct{}{}
	dp[l][len(c)]++
	dp[l][len(c)] %= m

	_, _ = fmt.Fprintf(writer, "%d", dp[l][k])
}

func answer2(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	buf := make([]byte, bufferSize)
	sc.Buffer(buf, 1e+7)
	sc.Split(bufio.ScanWords)
	n := nextString()
	l := len(n)
	k := uint64(nextInt())
	m := uint64(1e9 + 7)
	c := make([]bool, 16)
	ck := make([]int, 16)
	cnt := 0

	s2i := map[byte]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'A': 10,
		'B': 11,
		'C': 12,
		'D': 13,
		'E': 14,
		'F': 15,
	}

	dp := make([][]int, l+1)
	for i := range dp {
		dp[i] = make([]int, 17)
	}

	dp[0][0] = 0
	for i := 0; i < l; i++ {
		if i > 0 {
			dp[i+1][1] += 15
		}
		for j := 1; j < 17; j++ {
			dp[i+1][j] += (dp[i][j] * j) % int(m)
			if j < 16 {
				dp[i+1][j+1] += dp[i][j] * (16 - j)
			}
		}
		debugf("%d, %d, %v\n", i, s2i[n[i]], dp)

		for j := s2i[n[i]] - 1; j >= 0; j-- {
			if c[j] == false {
				if i > 0 || j > 0 {
					dp[i+1][cnt+1]++
				}
			} else {
				dp[i+1][cnt]++
			}
		}

		if c[s2i[n[i]]] == false {
			cnt++
			c[s2i[n[i]]] = true
			ck[i] = cnt
		}
		debugf("%d, %v\n", i, dp)

	}
	if cnt <= int(k) {
		dp[l][cnt]++
	}
	// dp[i+1][j] = dp[i][j] * j
	// dp[i+1][j+1] = dp[i][j] * (16-j)
	// dp[i+1][1] += 15

	debugf("%v\n\n", dp)
	_, _ = fmt.Fprintf(writer, "%d", dp[l][k])
}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

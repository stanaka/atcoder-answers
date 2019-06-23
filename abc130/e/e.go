package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

const BufferSize = 1024

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

var production bool

func debugf(format string, v ...interface{}) {
	if !production {
		fmt.Printf(format, v...)
	}
}

const MOD = 1000000000 + 7

// sum
// 1 2 3
// 2 5 9
// 3 9 19

// dp
// 1 1 1
// 1 2 3
// 1 3 6

func answer(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	buf := make([]byte, BufferSize)
	sc.Buffer(buf, 1e+6)
	sc.Split(bufio.ScanWords)
	n := nextInt()
	m := nextInt()
	s := make([]int, n)
	t := make([]int, m)
	dp := make([][]int, n)
	sum := make([][]int, n)

	for i := range s {
		s[i] = nextInt()
		dp[i] = make([]int, m)
		sum[i] = make([]int, m)
	}
	for i := range t {
		t[i] = nextInt()
	}

	for i := range s {
		for j := range t {
			if s[i] == t[j] {
				dp[i][j] = 1
				if i > 0 && j > 0 {
					dp[i][j] += sum[i-1][j-1]
					dp[i][j] %= MOD
				}
			}
			sum[i][j] = dp[i][j]
			if i > 0 {
				sum[i][j] += sum[i-1][j]
				sum[i][j] %= MOD
			}
			if j > 0 {
				sum[i][j] += sum[i][j-1]
				sum[i][j] %= MOD
			}
			if i > 0 && j > 0 {
				sum[i][j] -= sum[i-1][j-1]
				if sum[i][j] < 0 {
					sum[i][j] += MOD
				}
				sum[i][j] %= MOD
			}
		}
	}
	debugf("dp: %v\n", dp)
	debugf("sum: %v\n", sum)
	_, _ = fmt.Fprintf(writer, "%d", sum[n-1][m-1]+1)
}

func seek(s *[]int, t *[]int, c *[][]int, i int, j int) int {
	ans := 0
	if i+1 >= len(*s) || j+1 >= len(*t) {
		return 0
	}
	if (*c)[i+1][j+1] >= 0 {
		return (*c)[i+1][j+1]
	}

	for k := i + 1; k < len(*s); k++ {
		for l := j + 1; l < len(*t); l++ {
			if (*s)[k] == (*t)[l] {
				ans++
				ans %= MOD
				ans += seek(s, t, c, k, l)
				ans %= MOD
			}
		}
	}

	(*c)[i+1][j+1] = ans
	//debugf("i: %d, j: %d, ans: %d\n", i, j, ans)
	return ans
}

func answer_naive(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	buf := make([]byte, BufferSize)
	sc.Buffer(buf, 1e+6)
	sc.Split(bufio.ScanWords)
	n := nextInt()
	m := nextInt()
	s := make([]int, n)
	t := make([]int, m)
	c := make([][]int, n)

	for i := range s {
		s[i] = nextInt()
		c[i] = make([]int, m)
		for j := range t {
			c[i][j] = -1
		}
	}
	for i := range t {
		t[i] = nextInt()
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			_ = seek(&s, &t, &c, n-i-1, n-j-1)
		}
	}

	ans := 1
	ans += seek(&s, &t, &c, -1, -1)
	/*
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if s[i] == t[j] {
					ans++
					ans %= MOD
					ans += seek(&s, &t, &c, i, j)
					ans %= MOD
				}
			}
		}
	*/
	_, _ = fmt.Fprintf(writer, "%d", ans)
}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

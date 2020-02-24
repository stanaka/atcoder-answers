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

var production bool

func debugf(format string, v ...interface{}) {
	if !production {
		fmt.Printf(format, v...)
	}
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
	sc = bufio.NewScanner(reader)
	buf := make([]byte, bufferSize)
	sc.Buffer(buf, 1e+7)
	sc.Split(bufio.ScanWords)
	n := nextInt()
	k := nextInt()

	var c combi
	mod := int(1e9) + 7
	c.init(n+1, mod)
	ans := 0
	if n < k {
		k = n - 1
	}
	for i := 0; i <= k; i++ {
		ans += c.combi(n, i) * c.combi(n-1, n-i-1)
		ans %= mod
		//debugf("i: %d, ans: %d, diff:%d\n", i, ans, c.combi(n, i)*c.combi(n-1, n-i-1))
	}

	_, _ = fmt.Fprintf(writer, "%d", ans)
}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

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
	buf := make([]byte, BufferSize)
	sc.Buffer(buf, 1e+6)
	sc.Split(bufio.ScanWords)
	r1 := nextInt()
	c1 := nextInt()
	r2 := nextInt()
	c2 := nextInt()

	var c combi
	m := pow(10, 9) + 7
	c.init(r2+c2+4, m)
	ans := 0
	/*
		for i := r1; i < r2+1; i++ {
			for j := c1; j < c2+1; j++ {
				ans += c.combi(i+j, i)
				ans %= m
			}
		}
	*/

	ans += c.combi(r2+c2+2, r2+1)
	ans %= m
	ans -= c.combi(r2+c1+1, r2+1)
	if ans < 0 {
		ans += m
	}
	ans %= m
	ans -= c.combi(r1+c2+1, r1)
	if ans < 0 {
		ans += m
	}
	ans %= m
	ans += c.combi(r1+c1, r1)
	ans %= m

	_, _ = fmt.Fprintf(writer, "%d", ans)
}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

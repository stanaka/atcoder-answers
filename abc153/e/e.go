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

type magic struct {
	a int
	b int
}

func answer(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	buf := make([]byte, BufferSize)
	sc.Buffer(buf, 1e+6)
	sc.Split(bufio.ScanWords)
	h := nextInt()
	n := nextInt()
	m := make([]magic, n)

	for i := range m {
		m[i].a = nextInt()
		m[i].b = nextInt()
	}

	dp := make([]int, h+1)
	for i := 1; i <= h; i++ {
		minCost := -1
		for j := range m {
			cost := m[j].b
			if i-m[j].a >= 0 {
				cost += dp[i-m[j].a]
			}
			if minCost == -1 || minCost > cost {
				minCost = cost
			}
		}
		dp[i] = minCost
	}

	_, _ = fmt.Fprintf(writer, "%d", dp[h])
}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

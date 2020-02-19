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

func nextString() string {
	if !sc.Scan() {
		panic(nil)
	}
	return sc.Text()
}

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

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func solve(s []int) int {
	dp0 := make([]int, len(s))
	dp1 := make([]int, len(s))
	for i := range s {
		dp0[i] = s[i]
		dp1[i] = 10 - s[i]
		if i > 0 {
			dp0[i] += min(dp0[i-1], dp1[i-1])
			dp1[i] += min(dp0[i-1]+1, dp1[i-1]-1)
		} else {
			dp1[i]++
		}
	}
	return min(dp0[len(s)-1], dp1[len(s)-1])
}

func answer(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	buf := make([]byte, bufferSize)
	sc.Buffer(buf, 1e+7)
	sc.Split(bufio.ScanWords)
	n := nextString()
	//debugf("%v\n", n)
	//s := make([]int, len(n)+1)
	s := make([]int, len(n))
	for i := 0; i < len(n); i++ {
		d := int(n[i] - '0')
		s[i] = d
	}
	ans := solve(s)
	//debugf("%v\n", s)
	_, _ = fmt.Fprintf(writer, "%d", ans)
}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

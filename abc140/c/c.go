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

func min(a ...int) int {
	res := a[0]
	for i := 1; i < len(a); i++ {
		if res > a[i] {
			res = a[i]
		}
	}
	return res
}

func answer(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	buf := make([]byte, BufferSize)
	sc.Buffer(buf, 1e+6)
	sc.Split(bufio.ScanWords)
	n := nextInt()
	s := make([]int, n+1)
	for i := 1; i < n; i++ {
		s[i] = nextInt()
	}
	s[0] = s[1]
	s[n] = s[n-1]

	ans := 0
	for i := 0; i < n; i++ {
		ans += min(s[i], s[i+1])

	}

	_, _ = fmt.Fprintf(writer, "%d", ans)
}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

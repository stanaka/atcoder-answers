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

func answer(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	buf := make([]byte, BufferSize)
	sc.Buffer(buf, 1e+6)
	sc.Split(bufio.ScanWords)
	n := nextInt()
	a := make([]int, n)
	b := make([]int, n)
	c := make([]int, n-1)

	for i := range a {
		a[i] = nextInt() - 1
	}
	for i := range b {
		b[i] = nextInt()
	}
	for i := range c {
		c[i] = nextInt()
	}
	ans := 0
	for i := range a {
		debugf("a: %d:%d, %d\n", a[i], b[a[i]], ans)
		ans += b[a[i]]
		if i > 0 && a[i-1]+1 == a[i] {
			debugf("b: %d:%d, %d\n", a[i-1], c[a[i-1]], ans)
			ans += c[a[i-1]]
		}
	}

	_, _ = fmt.Fprintf(writer, "%d", ans)
}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

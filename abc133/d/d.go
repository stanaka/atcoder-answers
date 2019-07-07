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

/*
3-8+7-5+5
2
8-7+5-5+3
4
7-5+5-3+8
12
5-5+3-8+7
2
5-3+8-7+5
8
*/

func answer(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	buf := make([]byte, BufferSize)
	sc.Buffer(buf, 1e+6)
	sc.Split(bufio.ScanWords)
	n := nextInt()
	a := make([]int, n)
	x := make([]int, n)

	for i := range a {
		a[i] = nextInt()
	}

	for i := 0; i < n; i++ {
		flag := i % 2
		if flag == 0 {
			x[0] += a[i]
		} else {
			x[0] -= a[i]
		}
	}
	for i := 1; i < n; i++ {
		x[i] = -x[i-1] + a[i-1]*2
	}

	/*
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				idx := (i + j) % n
				flag := j % 2
				if flag == 0 {
					x[i] += a[idx]
				} else {
					x[i] -= a[idx]
				}
			}
		}
	*/
	debugf("%v\n", x)

	for i := 0; i < n; i++ {
		_, _ = fmt.Fprintf(writer, "%d", x[i])
		if i < n-1 {
			_, _ = fmt.Fprint(writer, " ")
		}
	}
}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

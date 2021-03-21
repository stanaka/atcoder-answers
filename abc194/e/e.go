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

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

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
	n := nextInt()
	m := nextInt()
	a := make([]int, n)
	c := make([]int, n)

	min := 0
	for i := range a {
		a[i] = nextInt()
		if i < m {
			c[a[i]]++
		}
	}
	debugf("%d\n", a)
	for i := 0; i < m; i++ {
		if c[i] > 0 {
			min = i + 1
			debugf("%d\n", min)
		} else {
			break
		}
	}

	for i := 0; i < n-m; i++ {
		c[a[i]]--
		c[a[i+m]]++
		if c[a[i]] == 0 {
			if a[i] < min {
				min = a[i]
			}
		}
	}

	_, _ = fmt.Fprintf(writer, "%d", min)
}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

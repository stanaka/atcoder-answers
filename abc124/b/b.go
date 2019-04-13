package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func answer(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)
	n := nextInt()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}
	b := 0
	c := 0
	for i := 0; i < n; i++ {
		if b <= a[i] {
			b = a[i]
			c++
		}
	}
	_, _ = fmt.Fprintf(writer, "%d\n", c)
}

func main() {
	answer(os.Stdin, os.Stdout)
}

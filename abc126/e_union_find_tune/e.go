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

func ufRoot(e []int, x int) int {
	i := x
	for {
		if e[i] == i {
			e[x] = i
			return i
		}
		i = e[i]
	}
}

func ufUnite(e []int, x int, y int) {
	rx := ufRoot(e, x)
	ry := ufRoot(e, y)
	if rx != ry {
		e[rx] = ry
		e[x] = ry
	}
}

func answer(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	m := nextInt()
	t := make([]int, n)
	for i := 0; i < n; i++ {
		t[i] = i
	}

	for i := 0; i < m; i++ {
		x := nextInt()
		y := nextInt()
		_ = nextInt()
		x--
		y--
		ufUnite(t, x, y)
	}
	//fmt.Printf("%v\n", t)
	count := 0
	for i := 0; i < n; i++ {
		if t[i] == i {
			count++
		}
	}
	_, _ = fmt.Fprintf(writer, "%d", count)
}

func main() {
	answer(os.Stdin, os.Stdout)
}

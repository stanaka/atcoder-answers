package main

import (
	"fmt"
	"io"
	"os"
)

func ufRoot(e []int, x int) int {
	if e[x] == x {
		return x
	}
	return ufRoot(e, e[x])
}

func ufUnite(e []int, x int, y int) {
	rx := ufRoot(e, x)
	ry := ufRoot(e, y)
	if rx != ry {
		e[rx] = ry
		e[y] = ry
	}
}

func answer(reader io.Reader, writer io.Writer) {
	var n, m int
	_, _ = fmt.Fscan(reader, &n, &m)
	t := make([]int, n)
	for i := 0; i < n; i++ {
		t[i] = i
	}

	for i := 0; i < m; i++ {
		var x, y, z int
		_, _ = fmt.Fscan(reader, &x, &y, &z)
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

package main

import (
	"fmt"
	"io"
	"os"
)

func dfs(e [][]int, f []bool, i int) {
	f[i] = true
	for _, v := range e[i] {
		if f[v] != true {
			dfs(e, f, v)
		}
	}
}

func answer(reader io.Reader, writer io.Writer) {
	var n, m int
	_, _ = fmt.Fscan(reader, &n, &m)
	e := make([][]int, n)

	for i := 0; i < m; i++ {
		var x, y, z int
		_, _ = fmt.Fscan(reader, &x, &y, &z)
		x--
		y--
		e[x] = append(e[x], y)
		e[y] = append(e[y], x)
	}
	f := make([]bool, n)
	count := 0
	for i := 0; i < n; i++ {
		if f[i] == false {
			count++
			dfs(e, f, i)
		}
	}
	_, _ = fmt.Fprintf(writer, "%d", count)
}

func main() {
	answer(os.Stdin, os.Stdout)
}

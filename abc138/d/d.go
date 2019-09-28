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
	q := nextInt()

	e := make([][]int, n)
	s := make([]int, n)

	for i := 0; i < n-1; i++ {
		u := nextInt()
		v := nextInt()
		e[u-1] = append(e[u-1], v-1)
		e[v-1] = append(e[v-1], u-1)
	}
	for i := 0; i < q; i++ {
		p := nextInt()
		x := nextInt()
		s[p-1] += x
	}

	ans := make([]int, n)
	ok := make([]bool, n)
	var dfs func(int, int)
	dfs = func(v int, x int) {
		ans[v] += x + s[v]
		ok[v] = true
		for _, u := range e[v] {
			if !ok[u] {
				dfs(u, s[v]+x)
			}
		}
	}
	dfs(0, 0)

	for i := 0; i < n; i++ {
		_, _ = fmt.Fprintf(writer, "%d", ans[i])
		if i < n-1 {
			_, _ = fmt.Fprint(writer, " ")
		}
	}
}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

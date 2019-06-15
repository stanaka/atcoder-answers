package main

import (
	"bufio"
	"io"
)

type edge struct {
	node   int
	weight int
}

func answer(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	buf := make([]byte, BufferSize)
	sc.Buffer(buf, 1e+6)
	n := nextInt()
	e := make([][]edge, n)

	for i := 0; i < n-1; i++ {
		u := nextInt()
		v := nextInt()
		w := nextInt()
		e[u-1] = append(e[u-1], edge{v - 1, w})
		e[v-1] = append(e[v-1], edge{u - 1, w})
	}
}

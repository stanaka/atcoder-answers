package main

import (
	"fmt"
	"io"
	"os"
)

func answer(reader io.Reader, writer io.Writer) {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	m := uint64(1000000000 + 1)
	s := uint64(0)
	c := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		v := a[i]
		if v < 0 {
			v = -v
			c++
		}
		if uint64(v) < m {
			m = uint64(v)
		}
		s += uint64(v)
	}
	if c%2 == 0 {
		_, _ = fmt.Fprintf(writer, "%d", s)
	} else {
		_, _ = fmt.Fprintf(writer, "%d", s-m-m)
	}
}

func main() {
	answer(os.Stdin, os.Stdout)
}

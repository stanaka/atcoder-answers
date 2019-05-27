package main

import (
	"fmt"
	"io"
	"os"
)

func answer(reader io.Reader, writer io.Writer) {
	var a, b int
	_, _ = fmt.Fscan(reader, &a, &b)
	min := 1
	max := a

	for i:=0; i< b; i++ {
		var l, r int
		_, _ = fmt.Fscan(reader, &l, &r)
		if min < l {
			min = l
		}
		if r < max {
			max = r
		}
	}
	d := max - min + 1
	if d < 0 {
		d = 0
	}
	_, _ = fmt.Fprintf(writer, "%d", d)
}

func main() {
	answer(os.Stdin, os.Stdout)
}

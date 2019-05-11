package main

import (
	"fmt"
	"io"
	"os"
)

func answer(reader io.Reader, writer io.Writer) {
	var n, k int
	var s string
	_, _ = fmt.Fscan(reader, &n, &s, &k)
	var r string
	for i := 0; i < len(s); i++ {
		if s[i] == s[k-1] {
			r = r + string(s[k-1])
		} else {
			r = r + "*"
		}
	}
	_, _ = fmt.Fprintf(writer, "%s", r)
}

func main() {
	answer(os.Stdin, os.Stdout)
}

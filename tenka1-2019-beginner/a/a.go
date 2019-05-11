package main

import (
	"fmt"
	"io"
	"os"
)

func answer(reader io.Reader, writer io.Writer) {
	var a, b, c int
	_, _ = fmt.Fscan(reader, &a, &b, &c)
	if (a < c && c < b) || (b < c && c < a) {
		_, _ = fmt.Fprintf(writer, "Yes")
	} else {
		_, _ = fmt.Fprintf(writer, "No")
	}
}

func main() {
	answer(os.Stdin, os.Stdout)
}

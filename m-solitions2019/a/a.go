package main

import (
	"fmt"
	"io"
	"os"
)

func answer(reader io.Reader, writer io.Writer) {
	var a int
	_, _ = fmt.Fscan(reader, &a)
	n := 180 * (a-2)
	_, _ = fmt.Fprintf(writer, "%d", n)
}

func main() {
	answer(os.Stdin, os.Stdout)
}

package main

import (
	"fmt"
	"io"
	"os"
)

func answer(reader io.Reader, writer io.Writer) {
	var a, b int
	_, _ = fmt.Fscan(reader, &a)
	_, _ = fmt.Fscan(reader, &b)
	_, _ = fmt.Fprintf(writer, "%d", a-b+1)
}

func main() {
	answer(os.Stdin, os.Stdout)
}

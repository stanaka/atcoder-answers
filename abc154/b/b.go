package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

var sc *bufio.Scanner

const BufferSize = 1024

func nextString() string {
	if !sc.Scan() {
		panic(nil)
	}
	return sc.Text()
}

func answer(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	buf := make([]byte, BufferSize)
	sc.Buffer(buf, 1e+6)
	sc.Split(bufio.ScanWords)
	s := nextString()
	var ans bytes.Buffer

	for _ = range s {
		ans.WriteString("x")
	}

	_, _ = fmt.Fprintf(writer, "%s", ans.String())
}

func main() {
	answer(os.Stdin, os.Stdout)
}

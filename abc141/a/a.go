package main

import (
	"bufio"
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
	s := nextString()
	if s == "Sunny" {
		s = "Cloudy"
	} else if s == "Cloudy" {
		s = "Rainy"
	} else {
		s = "Sunny"
	}

	_, _ = fmt.Fprintf(writer, "%s", s)
}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

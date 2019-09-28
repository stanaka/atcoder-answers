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

	f := true
	for i, c := range s {
		if i%2 == 0 {
			if string(c) != "R" && string(c) != "U" && string(c) != "D" {
				f = false
			}
		} else {
			if string(c) != "L" && string(c) != "U" && string(c) != "D" {
				f = false
			}
		}
	}

	if f {
		_, _ = fmt.Fprint(writer, "Yes")
	} else {
		_, _ = fmt.Fprint(writer, "No")
	}
}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var sc *bufio.Scanner

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func answer(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)
	a := nextString()
	var c int
	for i := 0; i < len(a); i++ {
		if a[i] == '1' {
			c++
		}
	}
	fmt.Fprintf(writer, "%d", c)
}

func main() {
	answer(os.Stdin, os.Stdout)
}

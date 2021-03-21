package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

const bufferSize = 1024

func nextInt() int {
	if !sc.Scan() {
		panic(nil)
	}
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

var production bool

func debugf(format string, v ...interface{}) {
	if !production {
		fmt.Printf(format, v...)
	}
}

func answer(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	buf := make([]byte, bufferSize)
	sc.Buffer(buf, 1e+7)
	sc.Split(bufio.ScanWords)
	a := nextInt()
	b := nextInt()

	s := 4
	if a+b >= 15 && b >= 8 {
		s = 1
	} else if a+b >= 10 && b >= 3 {
		s = 2
	} else if a+b >= 3 {
		s = 3
	}

	_, _ = fmt.Fprintf(writer, "%d", s)
}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

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
	x := nextInt()
	a := x / 500
	x -= a * 500
	c := a * 1000
	debugf("%d\n", c)
	a = x / 5
	x -= a * 5
	c += a * 5
	debugf("%d\n", c)

	_, _ = fmt.Fprintf(writer, "%d", c)
}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

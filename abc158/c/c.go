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

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func answer(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	buf := make([]byte, bufferSize)
	sc.Buffer(buf, 1e+7)
	sc.Split(bufio.ScanWords)
	a := nextInt()
	b := nextInt()
	minA := (a*100 + 7) / 8
	maxA := ((a+1)*100+7)/8 - 1
	minB := (b*100 + 9) / 10
	maxB := ((b+1)*100+9)/10 - 1

	debugf("A: %d - %d, B: %d - %d\n", minA, maxA, minB, maxB)

	if max(minA, minB) > min(maxA, maxB) {
		_, _ = fmt.Fprintf(writer, "%d", -1)
	} else {
		_, _ = fmt.Fprintf(writer, "%d", max(minA, minB))
	}

}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

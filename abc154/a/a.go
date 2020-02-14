package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

const BufferSize = 1024

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
	_ = nextString()
	a := nextInt()
	b := nextInt()
	u := nextString()

	if s == u {
		a--
	} else {
		b--
	}

	_, _ = fmt.Fprintf(writer, "%d %d", a, b)
}

func main() {
	answer(os.Stdin, os.Stdout)
}

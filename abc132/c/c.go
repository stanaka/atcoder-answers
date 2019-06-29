package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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
	n := nextInt()
	s := make([]int, n)

	for i := range s {
		s[i] = nextInt()
	}
	sort.Ints(s)

	ceiling := n / 2
	floor := n/2 - 1
	diff := s[ceiling] - s[floor]

	_, _ = fmt.Fprintf(writer, "%d", diff)
}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

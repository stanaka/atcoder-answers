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
	k := nextInt()
	q := nextInt()
	s := make([]int, n)

	for i := range s {
		s[i] = k - q
	}

	for i := 0; i < q; i++ {
		s[nextInt()-1]++
	}

	for i := range s {
		if s[i] > 0 {
			_, _ = fmt.Fprint(writer, "Yes\n")
		} else {
			_, _ = fmt.Fprint(writer, "No\n")
		}
	}

}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

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
	p := make([]int, n)
	q := make([]int, n)

	for i := range p {
		p[i] = nextInt()
		q[i] = p[i]
	}
	sort.Ints(q)
	count := 0
	for i := range p {
		if p[i] != q[i] {
			count++
		}
	}

	if count <= 2 {
		_, _ = fmt.Fprint(writer, "YES")
	} else {
		_, _ = fmt.Fprint(writer, "NO")
	}
}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

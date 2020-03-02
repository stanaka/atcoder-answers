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
	n := nextInt()
	m := nextInt()
	d := make([]int, n)

	for i := range d {
		d[i] = -1
	}
	for i := 0; i < m; i++ {
		s := nextInt() - 1
		c := nextInt()
		if d[s] >= 0 && d[s] != c {
			_, _ = fmt.Fprint(writer, "-1")
			return
		}
		if n > 1 && s == 0 && c == 0 {
			_, _ = fmt.Fprint(writer, "-1")
			return
		}
		d[s] = c
	}

	for i := range d {
		if i == 0 && n > 1 {
			if d[i] > 0 {
				_, _ = fmt.Fprintf(writer, "%d", d[i])
			} else {
				_, _ = fmt.Fprintf(writer, "%d", 1)
			}
		} else {
			if d[i] > 0 {
				_, _ = fmt.Fprintf(writer, "%d", d[i])
			} else {
				_, _ = fmt.Fprintf(writer, "%d", 0)
			}
		}
	}
}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

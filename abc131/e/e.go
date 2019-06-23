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

	d := (n-1)*(n-2)/2 - k
	if d < 0 {
		_, _ = fmt.Fprint(writer, "-1")
		return
	}

	g := map[int]map[int]struct{}{}

	count := 0
	for i := 1; i < n; i++ {
		if _, ok := g[0]; !ok {
			g[0] = map[int]struct{}{}
		}
		g[0][i] = struct{}{}
		count++
	}

	idx1 := 1
	idx2 := idx1 + 1
	for d > 0 {
		if idx2 < n {
			if _, ok := g[idx1]; !ok {
				g[idx1] = map[int]struct{}{}
			}
			g[idx1][idx2] = struct{}{}
			count++
			idx2++
		} else {
			idx1++
			idx2 = idx1 + 1
			if _, ok := g[idx1]; !ok {
				g[idx1] = map[int]struct{}{}
			}
			g[idx1][idx2] = struct{}{}
			count++
			idx2++
		}
		d--
	}

	_, _ = fmt.Fprintf(writer, "%d\n", count)
	for i := range g {
		for j := range g[i] {
			_, _ = fmt.Fprintf(writer, "%d %d\n", i+1, j+1)
		}
	}
}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

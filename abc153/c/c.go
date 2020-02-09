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

func answer(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	buf := make([]byte, BufferSize)
	sc.Buffer(buf, 1e+6)
	sc.Split(bufio.ScanWords)
	n := nextInt()
	k := nextInt()
	h := make([]int, n)

	for i := range h {
		h[i] = nextInt()
	}
	sort.Sort(sort.Reverse(sort.IntSlice(h)))

	ans := 0
	for i := range h {
		if k > 0 {
			k--
		} else {
			ans += h[i]
		}
	}

	_, _ = fmt.Fprintf(writer, "%d", ans)
}

func main() {
	answer(os.Stdin, os.Stdout)
}

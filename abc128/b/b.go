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

func nextString() string {
	if !sc.Scan() {
		panic(nil)
	}
	return sc.Text()
}

type re struct {
	index int
	name  string
	score int
}

type res []re

func (b res) Len() int {
	return len(b)
}

func (b res) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b res) Less(i, j int) bool {
	if b[i].name != b[j].name {
		return b[i].name < b[j].name
	} else {
		return b[i].score > b[j].score
	}
}

func answer(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	buf := make([]byte, BufferSize)
	sc.Buffer(buf, 1e+6)
	sc.Split(bufio.ScanWords)
	n := nextInt()
	r := make([]re, n)
	for i := 0; i < n; i++ {
		r[i].index = i + 1
		r[i].name = nextString()
		r[i].score = nextInt()
	}
	sort.Sort(res(r))

	for i := 0; i < n; i++ {
		_, _ = fmt.Fprintf(writer, "%d", r[i].index)
		if i != n-1 {
			_, _ = fmt.Fprint(writer, "\n")
		}
	}
}

func main() {
	answer(os.Stdin, os.Stdout)
}

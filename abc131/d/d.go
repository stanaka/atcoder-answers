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

type task struct {
	a int
	b int
}

type tasks []task

func (b tasks) Len() int           { return len(b) }
func (b tasks) Less(i, j int) bool { return b[i].b < b[j].b }
func (b tasks) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func answer(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	buf := make([]byte, BufferSize)
	sc.Buffer(buf, 1e+6)
	sc.Split(bufio.ScanWords)
	n := nextInt()
	t := make([]task, n)

	for i := range t {
		t[i].a = nextInt()
		t[i].b = nextInt()
	}
	sort.Sort(sort.Reverse(tasks(t)))

	d := t[0].b

	for i := range t {
		if d > t[i].b {
			d = t[i].b
		}
		d -= t[i].a
		debugf("i: %d, a: %d, d: %d\n", i, t[i].a, d)
		if d < 0 {
			break
		}
	}

	if d >= 0 {
		_, _ = fmt.Fprint(writer, "Yes")
	} else {
		_, _ = fmt.Fprint(writer, "No")
	}
}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

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

type bc struct {
	num int
	t   int
	end int
	pos int
}

type bcs []bc

func (b bcs) Len() int {
	return len(b)
}

func (b bcs) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b bcs) Less(i, j int) bool {
	return b[i].t < b[j].t
}

type ec struct {
	end int
	pos int
}

type ecs []ec

func (h ecs) Len() int           { return len(h) }
func (h ecs) Less(i, j int) bool { return h[i].pos < h[j].pos }
func (h ecs) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

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
	q := nextInt()

	s := make([]bc, 0, n)
	for i := 0; i < n; i++ {
		b := nextInt()
		e := nextInt()
		x := nextInt()
		s = append(s, bc{i, b - x, e - x, x})
	}
	sort.Sort(bcs(s))

	c := make([]ec, 0, n)
	idxS := 0
	debugf("s: %v\n", s)
	for i := 0; i < q; i++ {
		d := nextInt()
		debugf("d[%d]: %v\n", i, d)

		flag := true
		for j := 0; j < len(c); j++ {
			debugf("Check[%d]: %v\n", j, c[j])
			if c[j].end <= d {
				debugf("Remove[%d]\n", j)
				c = c[1:]
				j--
				if flag {
					sort.Sort(ecs(c))
				}
				flag = false
			} else {
				break
			}
		}

		flag = false
		for idxS < n && s[idxS].t <= d {
			if d < s[idxS].end {
				debugf("Push[%d]: %v\n", idxS, s[idxS])
				c = append(c, ec{s[idxS].end, s[idxS].pos})
				if c[0].pos > s[idxS].pos {
					flag = true
				}
			} else {
				debugf("Skip[%d]: %v\n", idxS, s[idxS])
			}
			debugf("c: %v\n", c)
			idxS++
		}
		if flag {
			sort.Sort(ecs(c))
		}

		if len(c) > 0 {
			debugf("Out: %v\n", c[0].pos)
			_, _ = fmt.Fprintf(writer, "%v\n", c[0].pos)
		} else {
			debugf("Out: -1\n")
			_, _ = fmt.Fprintln(writer, "-1")
		}
	}
}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

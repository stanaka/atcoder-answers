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

type bc struct {
	x int
	h int
}

type bcs []bc

func (b bcs) Len() int           { return len(b) }
func (b bcs) Less(i, j int) bool { return b[i].x < b[j].x }
func (b bcs) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

func answer(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	buf := make([]byte, BufferSize)
	sc.Buffer(buf, 1e+6)
	sc.Split(bufio.ScanWords)
	n := nextInt()
	d := nextInt()
	a := nextInt()
	s := make([]bc, n)

	for i := range s {
		s[i].x = nextInt()
		s[i].h = nextInt()
	}
	sort.Sort(bcs(s))

	ans := 0

	bm := make([][]int, 0, n)

	admg := 0
	debugf("%v\n", s)
	for l := range s {
		if s[l].h <= 0 {
			continue
		}

		// proceed r
		//for r < n && s[r].x <= s[l].x+2*d {
		//	r++
		//}
		for len(bm) > 0 && bm[0][1] < s[l].x-2*d {
			admg -= bm[0][0]
			bm = bm[1:]
		}

		// drop bomb
		s[l].h -= admg
		if s[l].h > 0 {
			count := (s[l].h + a - 1) / a
			ans += count
			dmg := a * count
			admg += dmg
			s[l].h -= dmg
			bm = append(bm, []int{dmg, s[l].x})
			debugf("%d, %d, %v\n", l, count, s)
		} else {
			debugf("%d, %d, %v\n", l, 0, s)
		}
	}

	_, _ = fmt.Fprintf(writer, "%d", ans)
}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

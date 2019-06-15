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

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type bc struct {
	x int
	y int
}

type bcs []bc

func (b bcs) Len() int { return len(b) }
func (b bcs) Less(i, j int) bool {
	if b[i].x != b[j].x {
		return b[i].x < b[j].x
	} else {
		return b[i].y < b[j].y
	}
}
func (b bcs) Swap(i, j int) { b[i], b[j] = b[j], b[i] }

func answer(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	buf := make([]byte, BufferSize)
	sc.Buffer(buf, 1e+6)
	sc.Split(bufio.ScanWords)
	n := nextInt()
	x := make([]bc, n)
	d := map[int]map[int]int{}

	for i := 0; i < n; i++ {
		x[i].x = nextInt()
		x[i].y = nextInt()
	}
	sort.Sort(bcs(x))

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if i == j {
				continue
			}
			dx := x[j].x - x[i].x
			dy := x[j].y - x[i].y
			if d[dx] == nil {
				d[dx] = map[int]int{}
			}
			d[dx][dy]++
			debugf("dx:%d, dy:%d, count:%d\n", dx, dy, d[dx][dy])
		}
	}

	debugf("%v\n", d)
	m := 0
	for i := range d {
		for j := range d[i] {
			m = max(m, d[i][j])
		}
	}

	_, _ = fmt.Fprintf(writer, "%d", n-m)
}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

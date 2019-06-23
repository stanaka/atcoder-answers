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

type UnionFind struct {
	e []int
}

func (u UnionFind) Root(x int) int {
	if u.e[x] < 0 {
		return x
	}
	//return u.Root(u.e[x])
	u.e[x] = u.Root(u.e[x])
	return u.e[x]
}

func (u UnionFind) Unite(x int, y int) bool {
	rx := u.Root(x)
	ry := u.Root(y)
	if rx != ry {
		u.e[rx] += u.e[ry]
		u.e[ry] = rx
		return true
	} else {
		return false
	}
}

func (u UnionFind) Size(x int) int {
	return -u.Root(x)
}

func answer(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	buf := make([]byte, BufferSize)
	sc.Buffer(buf, 1e+6)
	sc.Split(bufio.ScanWords)
	p := int(1e5)
	//p := 5
	n := nextInt()
	x := make([]int, n)
	y := make([]int, n)
	u := UnionFind{}
	u.e = make([]int, p*2+n)
	for i := range u.e {
		u.e[i] = -1
	}
	//debugf("e: %v\n", u.e)

	for i := 0; i < n; i++ {
		x[i] = nextInt()
		y[i] = nextInt()
		x[i]--
		y[i]--
		debugf("n:%d, x:%d, y:%d\n", n, x[i], y[i])
		u.Unite(i, n+x[i])
		u.Unite(i, n+p+y[i])
		//debugf("e: %v\n", u.e)
	}
	//debugf("e: %v\n", u.e)

	x1 := make([]map[int]struct{}, n+p*2)
	y1 := make([]map[int]struct{}, n+p*2)

	for i := 0; i < n; i++ {
		//debugf("i:%d, root:%d, x:%d, y:%d\n", n, x[i], y[i])
		if x1[u.Root(i)] == nil {
			x1[u.Root(i)] = map[int]struct{}{}
		}
		x1[u.Root(i)][x[i]] = struct{}{}
		if y1[u.Root(i)] == nil {
			y1[u.Root(i)] = map[int]struct{}{}
		}
		y1[u.Root(i)][y[i]] = struct{}{}
	}
	ans := 0
	for i := 0; i < n+p*2; i++ {
		nx := len(x1[i])
		ny := len(y1[i])
		ans += nx * ny
		if nx > 0 {
			debugf("i:%d, nx:%d, ny:%d\n", i, len(x1[i]), len(y1[i]))
		}
	}
	ans -= n

	_, _ = fmt.Fprintf(writer, "%d", ans)
}

func answer2(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	buf := make([]byte, BufferSize)
	sc.Buffer(buf, 1e+6)
	sc.Split(bufio.ScanWords)
	n := nextInt()
	v := map[int]map[int]struct{}{}
	w := map[int]map[int]struct{}{}

	for i := 0; i < n; i++ {
		x := nextInt()
		y := nextInt()
		if _, ok := v[x]; !ok {
			v[x] = map[int]struct{}{}
		}
		v[x][y] = struct{}{}
		if _, ok := w[y]; !ok {
			w[y] = map[int]struct{}{}
		}
		w[y][x] = struct{}{}
	}

	debugf("v:%v, w:%v\n", v, w)

	ans := 0
	for len(v) > 0 {
		for x := range v {
			count := 0
			v1 := map[int]struct{}{x: {}}
			vSet := map[int]struct{}{x: {}}
			w1 := map[int]struct{}{}
			for len(vSet) > 0 {
				for x := range vSet {
					delete(vSet, x)
					for y := range v[x] {
						if _, ok := w1[y]; !ok {
							w1[y] = struct{}{}
							count++
							debugf("[%d], v1:%v, w1:%v\n", count, v1, w1)
						}
						for x1 := range w[y] {
							if _, ok := v1[x1]; !ok {
								v1[x1] = struct{}{}
								vSet[x1] = struct{}{}
								count++
								debugf("[%d], v1:%v, w1:%v\n", count, v1, w1)
							}
						}
					}
				}
			}
			ans += len(v1)*len(w1) - count
			debugf("v1:%v, w1:%v, ans:%d, count:%d\n", v1, w1, ans, count)
			for x1 := range v1 {
				delete(v, x1)
			}
			for y1 := range w1 {
				delete(w, y1)
			}
		}
	}

	_, _ = fmt.Fprintf(writer, "%d", ans)
}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func answer(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	buf := make([]byte, BufferSize)
	sc.Buffer(buf, 1e+6)
	sc.Split(bufio.ScanWords)
	n := nextInt()
	s := make([]int, n)
	x := make([]int, n)
	y := make([]int, n)

	for i := range s {
		s[i] = nextInt()
	}
	sort.Ints(s)
	p := 1
	for p < n && s[p] < 0 {
		p++
	}
	if p == n {
		p--
	}
	debugf("s:%v p:%d\n", s, p)
	idx := 0
	x[0] = s[0]
	for i := p; i < n-1; i++ {
		y[idx] = s[i]
		idx++
		x[idx] = x[idx-1] - y[idx-1]
		debugf("a: x[%d]:%d = x[%d]:%d - y[%d]:%d\n", idx, x[idx], idx-1, x[idx-1], idx-1, y[idx-1])
	}
	y[idx] = x[idx]
	x[idx] = s[len(s)-1]
	idx++
	x[idx] = x[idx-1] - y[idx-1]
	debugf("b: x[%d]:%d = x[%d]:%d - y[%d]:%d\n", idx, x[idx], idx-1, x[idx-1], idx-1, y[idx-1])

	for i := 0; i < p-1; i++ {
		y[idx] = s[i+1]
		idx++
		x[idx] = x[idx-1] - y[idx-1]
		debugf("c: x[%d]:%d = x[%d]:%d - y[%d]:%d\n", idx, x[idx], idx-1, x[idx-1], idx-1, y[idx-1])
	}
	debugf("x:%v\n", x)
	debugf("y:%v\n", y)
	_, _ = fmt.Fprintf(writer, "%d\n", x[idx])
	for i := 0; i < n-1; i++ {
		_, _ = fmt.Fprintf(writer, "%d %d\n", x[i], y[i])
	}
}

func answer2(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	buf := make([]byte, BufferSize)
	sc.Buffer(buf, 1e+6)
	sc.Split(bufio.ScanWords)
	n := nextInt()
	s := make([]int, n)
	p := make([]int, 0, n)
	m := make([]int, 0, n)

	for i := range s {
		s[i] = nextInt()
		if s[i] >= 0 {
			p = append(p, s[i])
		}
		if s[i] < 0 {
			m = append(m, s[i])
		}
	}
	sort.Ints(p)
	sort.Sort(sort.Reverse(sort.IntSlice(m)))
	//sort.Ints(m)
	debugf("s:%v\n", s)
	debugf("p:%v (%d)\n", p, len(p))
	debugf("m:%v (%d)\n", m, len(m))

	if len(p) > 0 && len(m) > 0 {
		ans := 0
		ip := 0
		im := 0
		for i := 0; i < n; i++ {
			ans += abs(s[i])
		}
		_, _ = fmt.Fprintf(writer, "%d\n", ans)
		for i := 0; i < n-1; i++ {
			var x, y int
			if ip < len(p)-1 {
				x = m[im]
				y = p[ip]
				ip++
				m[im] = x - y
			} else {
				x = p[ip]
				y = m[im]
				im++
				p[ip] = x - y
			}
			_, _ = fmt.Fprintf(writer, "%d %d\n", x, y)
		}
	}
	if len(p) > 0 && len(m) == 0 {
		ans := 0
		for i := 2; i < n; i++ {
			ans += s[i]
		}
		ans += s[1] - s[0]
		_, _ = fmt.Fprintf(writer, "%d\n", ans)
		if len(s) == 2 {
			s[0], s[1] = s[1], s[0]
		}
		_, _ = fmt.Fprintf(writer, "%d %d\n", s[0], s[1])
		s[1] = s[0] - s[1]
		for i := 2; i < n; i++ {
			var x, y int
			y = s[i]
			x = s[i-1]
			s[i] = s[i-1] - s[i]
			if i == n-1 {
				x, y = y, x
			}
			_, _ = fmt.Fprintf(writer, "%d %d\n", x, y)
		}

	}
	if len(p) == 0 && len(m) > 0 {
		ans := 0
		for i := 2; i < n; i++ {
			ans += -m[i]
		}
		ans += m[0] - m[1]
		_, _ = fmt.Fprintf(writer, "%d\n", ans)
		_, _ = fmt.Fprintf(writer, "%d %d\n", m[0], m[1])
		m[1] = m[0] - m[1]
		for i := 2; i < n; i++ {
			var x, y int
			y = m[i]
			x = m[i-1]
			m[i] = m[i-1] - m[i]
			_, _ = fmt.Fprintf(writer, "%d %d\n", x, y)
		}
	}

}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

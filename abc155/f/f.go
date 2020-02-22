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

type pair struct {
	pos int
	f   int
}

func lowerBound(a []pair, x pair) int {
	l := 0
	r := len(a)
	for l <= r {
		m := (l + r) / 2
		if len(a) == m {
			break
		}
		if a[m].pos >= x.pos || (a[m].pos == x.pos && a[m].f >= x.f) {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}

func upperBound(a []pair, x pair) int {
	l := 0
	r := len(a)
	for l <= r {
		m := (l + r) / 2
		if len(a) == m {
			break
		}
		if a[m].pos <= x.pos || (a[m].pos == x.pos && a[m].f <= x.f) {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return l
}

type pairs []pair

func (b pairs) Len() int { return len(b) }
func (b pairs) Less(i, j int) bool {
	if b[i].pos == b[j].pos {
		return b[i].f < b[j].f
	}
	return b[i].pos < b[j].pos
}
func (b pairs) Swap(i, j int) { b[i], b[j] = b[j], b[i] }

type solver struct {
	a    []pair
	x    []int
	g    [][]pair
	ans  []int
	used []bool
}

func (s *solver) dfs(v int) int {
	s.used[v] = true
	res := s.x[v]
	debugf("dfscall: v: %d, res:%d\n", v, res)
	for _, e := range s.g[v] {
		if s.used[e.pos] {
			continue
		}
		r := s.dfs(e.pos)
		if r == 1 {
			s.ans = append(s.ans, e.f)
		}
		res ^= r
	}
	debugf("dfs ret: v: %d, res:%d, ans:%v\n", v, res, s.ans)
	return res
}

func answer(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	buf := make([]byte, bufferSize)
	sc.Buffer(buf, 1e+7)
	sc.Split(bufio.ScanWords)
	n := nextInt()
	m := nextInt()
	s := solver{}
	s.a = make([]pair, n)
	s.x = make([]int, n+1)
	s.g = make([][]pair, n+1)
	s.used = make([]bool, n+1)

	for i := range s.a {
		s.a[i].pos = nextInt()
		s.a[i].f = nextInt()
	}
	sort.Sort(pairs(s.a))

	s.x[0] = s.a[0].f
	for i := 0; i < n-1; i++ {
		s.x[i+1] = s.a[i].f ^ s.a[i+1].f
	}
	s.x[n] = s.a[n-1].f

	for i := 0; i < m; i++ {
		l, r := nextInt(), nextInt()
		l = lowerBound(s.a, pair{l, 0})
		r = upperBound(s.a, pair{r, 1})
		s.g[l] = append(s.g[l], pair{r, i + 1})
		s.g[r] = append(s.g[r], pair{l, i + 1})
	}

	for i := 0; i <= n; i++ {
		if s.used[i] {
			continue
		}
		if s.dfs(i) == 1 {
			_, _ = fmt.Fprint(writer, "-1")
			return
		}
	}
	sort.Ints(s.ans)
	debugf("a:%v\n", s.a)
	debugf("x:%v\n", s.x)
	debugf("g:%v\n", s.g)
	debugf("ans:%v\n", s.ans)

	_, _ = fmt.Fprintf(writer, "%d\n", len(s.ans))
	for i, v := range s.ans {
		_, _ = fmt.Fprintf(writer, "%d", v)
		if i != len(s.ans)-1 {
			_, _ = fmt.Fprintf(writer, " ")
		}
	}
}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

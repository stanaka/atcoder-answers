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

type edge struct {
	src    int
	dst    int
	weight int
}

func reachable(g [][]int, start int) []bool {
	n := len(g)
	ok := make([]bool, n)
	var dfs func(int)
	dfs = func(v int) {
		ok[v] = true
		for _, u := range g[v] {
			if !ok[u] {
				dfs(u)
			}
		}
	}
	dfs(start)
	return ok
}

func answer(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	buf := make([]byte, BufferSize)
	sc.Buffer(buf, 1e+6)
	sc.Split(bufio.ScanWords)
	n := nextInt()
	m := nextInt()
	p := nextInt()
	e := make([]edge, m)
	gn := make([][]int, n)
	gr := make([][]int, n)
	dis := make([]int, n)
	pre := make([]int, n)

	for i := 0; i < m; i++ {
		u := nextInt()
		v := nextInt()
		w := nextInt()
		e[i] = edge{u - 1, v - 1, -(w - p)}
		gn[u-1] = append(gn[u-1], v-1)
		gr[v-1] = append(gr[v-1], u-1)
	}
	okn := reachable(gn, 0)
	okr := reachable(gr, n-1)
	var es []edge
	for i := range e {
		if okn[e[i].src] && okr[e[i].src] && okn[e[i].dst] && okr[e[i].dst] {
			es = append(es, e[i])
		}
	}
	for i := 0; i < n; i++ {
		dis[i] = int(1e8)
		pre[i] = 0
	}

	dis[0] = 0
	count := 0
	update := true
	for update {
		update = false
		for j := range es {
			if dis[es[j].src]+es[j].weight < dis[es[j].dst] {
				dis[es[j].dst] = dis[es[j].src] + es[j].weight
				pre[es[j].dst] = es[j].src
				update = true
			}
		}
		count++
		if count > n {
			break
		}
	}

	ans := 0
	if count > n {
		ans = -1
	} else if dis[n-1] < 0 {
		ans = -dis[n-1]
	} else if dis[n-1] == int(1e8) {
		ans = -1
	}
	_, _ = fmt.Fprintf(writer, "%d", ans)
}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

package main

import (
	"bufio"
	"io"
)

type edge struct {
	node   int
	weight int
}

func answer(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	buf := make([]byte, BufferSize)
	sc.Buffer(buf, 1e+6)
	n := nextInt()
	e := make([][]edge, n)

	for i := 0; i < n-1; i++ {
		u := nextInt()
		v := nextInt()
		w := nextInt()
		e[u-1] = append(e[u-1], edge{v - 1, w})
		e[v-1] = append(e[v-1], edge{u - 1, w})
	}

	g := map[int]map[int]struct{}{}
	for i := 1; i < n; i++ {
		if _, ok := g[0]; !ok {
			g[0] = map[int]struct{}{}
		}
		g[0][i] = struct{}{}
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

func bfs(e [][]int, s int) (int, int) {
	set := e[s]
	visited := map[int]struct{}{s: {}}
	last := s
	count := 0
	for len(set) > 0 {
		nextSet := []int{}
		debugf("bfs: set: %v\n", set)

		for i := range set {
			visited[set[i]] = struct{}{}
			last = set[i]

			for _, j := range e[set[i]] {
				if _, ok := visited[j]; !ok {
					nextSet = append(nextSet, j)
				}
			}
		}
		set = nextSet
		count++
	}
	return last, count
}

func radius(e [][]int) int {
	farthest, _ := bfs(e, 0)
	debugf("farthest: %v\n", farthest)

	_, count := bfs(e, farthest)
	return count
}

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

func answer(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	buf := make([]byte, BufferSize)
	sc.Buffer(buf, 1e+6)
	sc.Split(bufio.ScanWords)
	n := nextInt()
	e := make([][]int, n)

	for i := 0; i < n-1; i++ {
		u := nextInt()
		v := nextInt()
		e[u-1] = append(e[u-1], v-1)
		e[v-1] = append(e[v-1], u-1)
	}
	debugf("e: %v\n", e)

	ans := radius(e)
	debugf("radius: %v\n", ans)

	if ans%3 != 1 {
		_, _ = fmt.Fprint(writer, "First")
	} else {
		_, _ = fmt.Fprint(writer, "Second")
	}
}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

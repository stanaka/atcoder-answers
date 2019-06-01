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

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

type edge struct {
	node   int
	weight int
}

type count struct {
	node  int
	count int
}

func answer(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)
	n := nextInt()
	e := make([][]edge, n)
	w := make([]int, n)
	c := make([]int, n)
	v := make([]int, n)
	f := make([]bool, n)

	for i := 0; i < n-1; i++ {
		a := nextInt()
		b := nextInt()
		e[a-1] = append(e[a-1], edge{b - 1, 0})
		e[b-1] = append(e[b-1], edge{a - 1, 0})
	}
	for i := 0; i < n; i++ {
		v[i] = nextInt()
	}
	sort.Sort(sort.Reverse(sort.IntSlice(v)))

	idx := 0
	tail := 1
	f[0] = true
	for idx < n {
		w[c[idx]] = v[idx]
		for _, v := range e[c[idx]] {
			if f[v.node] == false {
				c[tail] = v.node
				f[v.node] = true
				tail++
			}
		}
		idx++
	}

	sum := 0
	for i := 0; i < n; i++ {
		for j := 0; j < len(e[i]); j++ {
			x := w[i]
			if w[e[i][j].node] < w[i] {
				x = w[e[i][j].node]
			}
			e[i][j].weight = x
			sum += x
		}
	}
	//fmt.Printf("%v\n", v)
	//fmt.Printf("%v\n", c)
	//fmt.Printf("%v\n", w)
	//fmt.Printf("%v\n", e)

	_, _ = fmt.Fprintf(writer, "%d\n", sum/2)
	for i := 0; i < n; i++ {
		_, _ = fmt.Fprintf(writer, "%d", w[i])
		if i != n-1 {
			_, _ = fmt.Fprintf(writer, " ")
		}
	}
}

func main() {
	answer(os.Stdin, os.Stdout)
}

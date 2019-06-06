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

func pow(x int, y int) int {
	//_, _ = fmt.Printf("pow: %v, %v\n", x,y)
	z := x
	if y == 0 {
		return 1
	}
	var i int
	for i = 1; i < y; i++ {
		z = z * x
	}
	return z
}

func answer(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	buf := make([]byte, BufferSize)
	sc.Buffer(buf, 1e+6)
	sc.Split(bufio.ScanWords)
	n := nextInt()
	m := nextInt()
	l := make([][]int, m)
	p := make([]int, m)
	for i := 0; i < m; i++ {
		k := nextInt()
		l[i] = make([]int, k)
		for j := 0; j < k; j++ {
			l[i][j] = nextInt() - 1
		}
	}
	for i := 0; i < m; i++ {
		p[i] = nextInt()
	}

	a := make([]int, 10)
	for i := 0; i < 10; i++ {
		a[i] = pow(2, i)
	}

	//fmt.Printf("%v\n", l)

	count := 0
	for i := 0; i < pow(2, n); i++ {
		lit := true
		for j := 0; j < m; j++ {
			c := 0
			for k := 0; k < len(l[j]); k++ {
				//fmt.Printf("i:%v, j:%v, k:%v, l:%v, a:%v\n", i, j, k, l[j][k], a[l[j][k]])
				if a[l[j][k]]&i != 0 {
					c++
				}
			}
			//fmt.Printf("i:%v, j:%v, c:%v\n", i, j, c)
			if c%2 != p[j] {
				//fmt.Print("not lit!\n")
				lit = false
				break
			}
		}
		if lit == true {
			//fmt.Print("lit!\n")
			count++
		}
	}
	//fmt.Printf("%v\n", count)

	_, _ = fmt.Fprintf(writer, "%d", count)
}

func main() {
	answer(os.Stdin, os.Stdout)
}

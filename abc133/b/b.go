package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func distance(x, y []int, d int) float64 {
	ret := 0
	for i := 0; i < d; i++ {
		di := abs(x[i] - y[i])
		ret += di * di
	}
	return math.Sqrt(float64(ret))
}

func answer(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	buf := make([]byte, BufferSize)
	sc.Buffer(buf, 1e+6)
	sc.Split(bufio.ScanWords)
	n := nextInt()
	d := nextInt()
	x := make([][]int, n)

	for i := range x {
		x[i] = make([]int, d)
		for j := range x[i] {
			x[i][j] = nextInt()
		}
	}

	ans := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			d2 := distance(x[i], x[j], d)
			if float64(int64(d2)) == d2 {
				ans++
			}
		}
	}

	_, _ = fmt.Fprintf(writer, "%d", ans)
}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

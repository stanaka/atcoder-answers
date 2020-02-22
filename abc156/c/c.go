package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

func answer(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	buf := make([]byte, bufferSize)
	sc.Buffer(buf, 1e+7)
	sc.Split(bufio.ScanWords)
	n := nextInt()
	s := make([]int, n)
	sum := 0
	for i := range s {
		s[i] = nextInt()
		sum += s[i]
	}
	p := sum / n
	p1 := float64(sum) / float64(n)
	if p1-float64(p) >= 0.5 {
		p++
	}
	ans := 0
	//debugf("p:%d, sum:%d, n:%d\n", p, sum, n)
	for i := range s {
		ans += (s[i] - p) * (s[i] - p)
	}

	_, _ = fmt.Fprintf(writer, "%d", ans)
}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

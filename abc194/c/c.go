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
	c := make([]int, 400+1)

	for i := range s {
		s[i] = nextInt()
		c[s[i]+200]++
	}

	ans := 0
	for i := 0; i < 200*2; i++ {
		for j := i + 1; j < 200*2+1; j++ {
			diff := (i - 200) - (j - 200)
			ans += diff * diff * c[i] * c[j]
		}
	}

	_, _ = fmt.Fprintf(writer, "%d", ans)
}

func answer_naive(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	buf := make([]byte, bufferSize)
	sc.Buffer(buf, 1e+7)
	sc.Split(bufio.ScanWords)
	n := nextInt()
	s := make([]int, n)

	for i := range s {
		s[i] = nextInt()
	}

	ans := 0
	for i := 0; i < n-1; i++ {
		for j := i; j < n; j++ {
			diff := s[i] - s[j]
			ans += diff * diff
		}
	}

	_, _ = fmt.Fprintf(writer, "%d", ans)
}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

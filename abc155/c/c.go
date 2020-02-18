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

func nextString() string {
	if !sc.Scan() {
		panic(nil)
	}
	return sc.Text()
}

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
	buf := make([]byte, BufferSize)
	sc.Buffer(buf, 1e+6)
	sc.Split(bufio.ScanWords)
	n := nextInt()
	s := map[string]int{}

	max := 0
	for i := 0; i < n; i++ {
		st := nextString()
		s[st]++
		if max < s[st] {
			max = s[st]
		}
	}

	a := make([]string, 0)

	for i := range s {
		if s[i] == max {
			a = append(a, i)
		}
	}

	sort.Strings(a)
	for i := range a {
		_, _ = fmt.Fprintf(writer, "%s\n", a[i])
	}
}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

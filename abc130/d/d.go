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

func nextUint64() uint64 {
	if !sc.Scan() {
		panic(nil)
	}
	i, e := strconv.ParseUint(sc.Text(), 0, 64)
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
	k := nextUint64()
	a := make([]uint64, n)

	for i := range a {
		a[i] = nextUint64()
	}
	ans := 0
	s := 0
	sum := uint64(0)
	e := s
	sum += a[e]
	for s < n {
		for e < n {
			if sum >= k {
				ans += n - e
				debugf("[1] s: %d, e: %d, sum: %d, ans: %d\n", s, e, sum, ans)
				break
			}
			debugf("[2] s: %d, e: %d, sum: %d, ans: %d\n", s, e, sum, ans)
			e++
			if e < n {
				sum += a[e]
			}
		}
		sum -= a[s]
		s++
		//if sum >= k {
		//	ans++
		//}
		debugf("[3] s: %d, e: %d, sum: %d, ans: %d\n", s, e, sum, ans)
	}

	_, _ = fmt.Fprintf(writer, "%d", ans)
}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

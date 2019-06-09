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

func nextString() string {
	if !sc.Scan() {
		panic(nil)
	}
	return sc.Text()
}

var production bool

func debugf(format string, v ...interface{}) {
	if !production {
		fmt.Printf(format, v...)
	}
}

// 1: 0-1
// 2: 0-2
//    0-1-2
// 3: 0-1-2-3
//    0-2-3
//    0-1-3
// 4: 0-1-2-3-4
//    0-1-3-4
//    0-2-3-4
//    0-2-4
//    0-1-2-4

func answer(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	buf := make([]byte, BufferSize)
	sc.Buffer(buf, 1e+6)
	sc.Split(bufio.ScanWords)
	n := nextInt() + 1
	m := nextInt()
	a := make([]int, m+1)
	dp := make([]int, n)

	for i := 0; i < m; i++ {
		a[i] = nextInt()
	}
	idx := 0

	for i := range dp {
		if i == 0 {
			dp[i] = 1
			continue
		}
		if i == 1 {
			if a[idx] == i {
				dp[i] = 0
				idx++
			} else {
				dp[i] = 1
			}
			continue
		}
		if a[idx] == i {
			dp[i] = 0
			idx++
		} else {
			dp[i] = (dp[i-1] + dp[i-2]) % 1000000007
		}
	}
	debugf("%v\n", dp)

	_, _ = fmt.Fprintf(writer, "%d", dp[n-1])
}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

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

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

var production bool

func debugf(format string, v ...interface{}) {
	if !production {
		//fmt.Printf(format, v...)
	}
}

func answer(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	buf := make([]byte, BufferSize)
	sc.Buffer(buf, 1e+6)
	sc.Split(bufio.ScanWords)
	n := nextInt()
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}

	// 16
	// (4,1) 4,3,7,6,10,9,13,12,16 (4,12)
	// (7,4) 7,3,10,6,13,9,16 (7, 9)
	// (10,7) 10,3,13,6,16 (10,6)
	// (13,10) 13,3,16 (13,3)
	// (16,13) 16
	//
	// n-1 = (a-b)*x+a -> n-a-1 = (a-b) * x
	// a!=2(a-b) -> a != 2b
	// a!=3(a-b) -> 2a != 3b

	// 0 -4 0 -99 31 14 -15 -39 43 18 0`
	// (6,4) 6,2,8,4,10
	// -15+0+43+31+0
	// (10,8) 10
	// (8,6) 8,2,10 (8,2)
	// (6,4) 6,2,8,4,10 (6,4)
	// (4,2) 4,2,6,4 (xx)
	ans := 0
	//dp := make([]int, n)
	u := make([]bool, n)
	for c := 1; c < n; c++ {
		a := n - 1
		debugf("a: %d, c: %d\n", a, c)
		sc := s[a]
		//ans = max(ans, sc)
		//dp[0] = s[a]
		//ans = max(ans, dp[0])
		//debugf("i: %d, a: %d, score: %d\n", 0, a, dp[0])
		debugf("i: %d, a: %d, score: %d\n", 0, a, sc)
		i := 1
		a -= c
		for a > c {
			u[i*c] = true
			//if i*c >= n-1-i*c {
			if u[n-1-i*c] {
				break
			}
			/*if a%c == 0 && i*c >= a {
				debugf("skip. a: %d, c: %d\n", a, c)
				//a -= c
				break
			}*/
			sc += s[n-1-c*i] + s[c*i]
			ans = max(ans, sc)
			debugf("i: %d, a: %d, score: %d\n", i, a, sc)
			//dp[i] = dp[i-1] + s[n-1-c*i] + s[c*i]
			//ans = max(ans, dp[i])
			//debugf("i: %d, a: %d, score: %d\n", i, a, dp[i])
			i++
			a -= c
		}
		for k := 0; k*c < n-1; k++ {
			u[k*c] = false
		}
	}
	_, _ = fmt.Fprintf(writer, "%d", ans)
}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

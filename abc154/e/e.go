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

// 314159
// 9*9*5C2 = 810
// 1*(1+9+9+9+9) = 37
// 2*(9+9+9+9+9) = 90
// 100
//       0 1 2 3
// dp0 0 1 0 0 0
//     1 0 0 9 18
//     2 0 0 9
//       0 1 2 3
// dp1 0 0
//     1 0 1 1 1
//     2 0 1 2 2
func solve(n string, k int) int {
	//var dp [102][4][2]int
	l := len(n)
	dp := make([][][]int, l+1)
	for i := range dp {
		dp[i] = make([][]int, 4)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
		}
	}
	dp[0][0][0] = 1
	for i := 0; i < l; i++ {
		for j := 0; j < 4; j++ {
			for p := 0; p < 2; p++ {
				f := int(n[i] - '0')
				for d := 0; d < 10; d++ {
					ni, nj, np := i+1, j, p
					if d > 0 {
						nj++
					}
					if nj > k {
						continue
					}
					if p == 0 {
						switch {
						case d < f:
							np = 1
						case d > f:
							continue
						}
					}
					dp[ni][nj][np] += dp[i][j][p]
				}
			}
		}
	}
	debugf("%v\n", dp)

	return dp[l][k][0] + dp[l][k][1]
}

func solve2(n string, k int) int {
	for n[0] == '0' {
		if len(n) == 1 {
			return -1
		}
		n = n[1:]
	}
	ans := 0
	if k > len(n) {
		return -1
	}
	f := int(n[0] - '0')
	switch k {
	case 1:
		ans += f + (len(n)-1)*9
	case 2:
		if f > 0 {
			ans += 9 * 9 * (len(n) - 1) * (len(n) - 2) / 2
		}
		if f > 1 {
			ans += (f - 1) * (9 * (len(n) - 1))
		}
	case 3:
		if f > 0 {
			ans += 9 * 9 * 9 * (len(n) - 1) * (len(n) - 2) * (len(n) - 3) / 6
		}
		if f > 1 {
			ans += (f - 1) * (9 * 9 * (len(n) - 1) * (len(n) - 2) / 2)
		}
	}
	if k > 1 {
		ret := solve(n[1:], k-1)
		if ret == -1 {
			return -1
		}
		ans += ret
	}
	return ans
}

func answer(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	buf := make([]byte, BufferSize)
	sc.Buffer(buf, 1e+6)
	sc.Split(bufio.ScanWords)
	n := nextString()
	k := nextInt()

	ans := solve(n, k)
	//if ans == -1 {
	//		ans = 0
	//	}
	_, _ = fmt.Fprintf(writer, "%d", ans)
}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

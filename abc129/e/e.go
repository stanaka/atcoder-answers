package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var sc *bufio.Scanner

const BufferSize = 1024

func nextString() string {
	if !sc.Scan() {
		panic(nil)
	}
	return sc.Text()
}

// 1 - 3 patterns
// 0 0,1
// 1 0
// 1*3

// 11 - 9 patterns
// 00 00,01,10,11
// 01 00,10
// 10 00,01
// 11 00
// 3*3
// 1*3*3
// [2,4] [1,5]

// 10 - 5 patterns
// 00 00,01,10
// 01 00
// 10 00
// 3 + 2
// 3(1)-5(10)
// [2,2] [1,3]

// 100 - 11 patterns
// 000 000,001,010,011,100
// 001 000,010
// 010 000,001
// 011 000
// 100 000
// 5(10)-11(100)
// [2,2,2] [1,3,9]

// 110 - 19 patterns
// 000 000,001,010,011,100,101,110
// 001 000,010,100
// 010 000,001,100
// 011 000
// 100 000,001,010
// 101 000
// 110 000
// 1*3*3
// 9 + 5 + 5
// 1*3*3 + 2*(1*3+2*1)
// 9 + 5 + 5
// 1*3*3 + 1*1*3 + 1*2*1 + 1*1*3 + 1*2*1
// 9(11)-19(110)
// [2,4,4] [1,5,15]

// 111 - 27 patterns
// [2,4,8] [1,5,19]

func answer(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	buf := make([]byte, BufferSize)
	sc.Buffer(buf, 1e+7)
	sc.Split(bufio.ScanWords)
	s := nextString()
	mod := 1000*1000*1000 + 7

	dp1 := make([]int, len(s))
	dp2 := make([]int, len(s))

	for i := 0; i < len(s); i++ {
		if i == 0 {
			dp1[0] = 2
			dp2[0] = 1
			continue
		}

		dp1[i] = dp1[i-1]
		dp2[i] = dp2[i-1] * 3

		if string(s[i]) == "1" {
			dp1[i] += dp1[i-1]
			dp2[i] += dp1[i-1]
		}
		dp1[i] %= mod
		dp2[i] %= mod
	}
	_, _ = fmt.Fprintf(writer, "%d", (dp1[(len(s)-1)]+dp2[len(s)-1])%mod)
}

func main() {
	answer(os.Stdin, os.Stdout)
}

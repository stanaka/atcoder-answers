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

var production bool

func debugf(format string, v ...interface{}) {
	if !production {
		fmt.Printf(format, v...)
	}
}

// ? .. 1
// ?? .. 8 (5,18,31,44,57,70,83,96)
// ??? ..

func answer(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	buf := make([]byte, BufferSize)
	sc.Buffer(buf, 1e+6)
	sc.Split(bufio.ScanWords)
	s := nextString()
	a := make([]int, 13)
	b := make([]int, 13)
	mod := int(1e9) + 7
	offset := 1
	// i == 0
	if string(s[len(s)-1]) == "?" {
		m := 0
		for j := 0; j < 10; j++ {
			a[m] = 1
			m += offset
			m %= 13
		}
	} else {
		n, _ := strconv.Atoi(string(s[len(s)-1]))
		m := (offset * n) % 13
		a[m%13] = 1
	}
	debugf("a[%d]: %v\n", 0, a)
	for i := 0; i < 13; i++ {
		b[i] = a[i]
		a[i] = 0
	}
	offset = (offset * 10) % 13

	// i> 0
	for i := 1; i < len(s); i++ {
		if string(s[len(s)-i-1]) == "?" {
			m := 0
			for j := 0; j < 10; j++ {
				for k := 0; k < 13; k++ {
					a[(k+m)%13] += b[k]
					a[(k+m)%13] %= mod
				}
				m += offset
				m %= 13
			}
		} else {
			n, _ := strconv.Atoi(string(s[len(s)-i-1]))
			m := (offset * n) % 13
			for k := 0; k < 13; k++ {
				a[(k+m)%13] = b[k]
			}
		}
		debugf("a[%d]: %v\n", i, a)
		for i := 0; i < 13; i++ {
			b[i] = a[i]
			a[i] = 0
		}
		offset = (offset * 10) % 13
	}
	debugf("b: %v\n", b)
	_, _ = fmt.Fprintf(writer, "%d", b[5])
}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

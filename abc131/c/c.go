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

func minUint64(a ...uint64) uint64 {
	res := a[0]
	for i := 1; i < len(a); i++ {
		if res > a[i] {
			res = a[i]
		}
	}
	return res
}

func GCD(a, b uint64) uint64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b uint64, integers ...uint64) uint64 {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func answer(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	buf := make([]byte, BufferSize)
	sc.Buffer(buf, 1e+6)
	sc.Split(bufio.ScanWords)
	a := nextUint64()
	b := nextUint64()
	c := nextUint64()
	d := nextUint64()
	ans := uint64(0)

	e := LCM(c, d)
	// 4 5 6 7 8 9
	// 6 9
	// 3 - 1
	cn := ((b / c) - ((a - 1) / c))
	// 4 6 8
	// 4 - 1
	dn := ((b / d) - ((a - 1) / d))
	// 6
	en := ((b / e) - ((a - 1) / e))
	//
	ans = (b - a + 1) - cn - dn + en
	debugf("ans: %d, cn: %d, dn: %d, en: %d (e: %d)\n", ans, cn, dn, en, e)

	_, _ = fmt.Fprintf(writer, "%d", ans)
}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

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

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

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

	mina := MaxInt
	seca := MaxInt
	inda := 0
	minb := MaxInt
	secb := MaxInt
	indb := 0
	minDiff := 0
	minSame := MaxInt
	for i := 0; i < n; i++ {
		a := nextInt()
		if a < mina {
			seca = mina
			mina = a
			inda = i
		} else if a < seca {
			seca = a
		}

		b := nextInt()
		if b < minb {
			secb = minb
			minb = b
			indb = i
		} else if b < secb {
			secb = b
		}
	}

	minSame = mina + minb
	if inda != indb {
		minDiff = max(mina, minb)
	} else {
		minDiff = min(max(mina, secb), max(seca, minb))
	}
	debugf("%d, %d\n", inda, indb)
	debugf("%d, %d\n", mina, minb)
	debugf("%d, %d\n", seca, secb)
	debugf("%d, %d\n", minDiff, minSame)

	_, _ = fmt.Fprintf(writer, "%d", min(minDiff, minSame))
}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

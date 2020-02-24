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

// 3,6,7,11,14
// 1,0,1,1,0
//  1,1,0,1

func answer(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	buf := make([]byte, bufferSize)
	sc.Buffer(buf, 1e+7)
	sc.Split(bufio.ScanWords)
	k := nextInt()
	q := nextInt()
	d := make([]int, k)
	d2 := make([]int, k)

	for i := range d {
		d[i] = nextInt()
	}
	for i := 0; i < q; i++ {
		n := nextInt()
		x := nextInt()
		m := nextInt()
		count := (n - 1) / k
		lastNum := (n - 1) % k
		numEq := 0
		cntEq := 0
		lastA := uint64(x)
		for j := range d {
			d2[j] = d[j] % m
			if d2[j] == 0 {
				numEq++
				if j < lastNum {
					cntEq++
				}
			} else {
				lastA += uint64(count) * uint64(d2[j])
				if j < lastNum {
					lastA += uint64(d2[j])
				}
			}
		}
		cntEq += numEq * count
		cntUp := lastA/uint64(m) - uint64(x)/uint64(m)
		ans := n - 1 - int(cntEq) - int(cntUp)
		//debugf("n:%d, x:%d, m:%d, count:%d, lastNum:%d, lastA:%d, cntEq:%d, cntUp:%d, ans:%d, d2:%v\n", n, x, m, count, lastNum, lastA, cntEq, cntUp, ans, d2)
		_, _ = fmt.Fprintf(writer, "%d", ans)
		if i != q-1 {
			_, _ = fmt.Fprint(writer, "\n")
		}
	}
}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

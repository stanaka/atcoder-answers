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

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func answer(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	buf := make([]byte, BufferSize)
	sc.Buffer(buf, 1e+6)
	sc.Split(bufio.ScanWords)
	h := nextInt()
	w := nextInt()
	s := make([]string, h)

	for i := range s {
		s[i] = nextString()
	}

	hs := make([][]int, h)
	for i := 0; i < h; i++ {
		hs[i] = make([]int, w)
		blockStart := 0
		for j := 0; j < w; j++ {
			if string(s[i][j]) == "#" {
				for k := blockStart; k < j; k++ {
					hs[i][k] = j - blockStart
				}
				blockStart = j + 1
				hs[i][j] = 0
			}
		}
		for k := blockStart; k < w; k++ {
			hs[i][k] = w - blockStart
		}
	}
	debugf("hs: %v\n", hs)

	ws := make([][]int, h)
	for i := 0; i < h; i++ {
		ws[i] = make([]int, w)
	}
	for i := 0; i < w; i++ {
		blockStart := 0
		for j := 0; j < h; j++ {
			if string(s[j][i]) == "#" {
				for k := blockStart; k < j; k++ {
					ws[k][i] = j - blockStart
				}
				blockStart = j + 1
				ws[j][i] = 0
			}
		}
		for k := blockStart; k < h; k++ {
			ws[k][i] = h - blockStart
		}
	}
	debugf("ws: %v\n", ws)

	ans := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			ans = max(ans, hs[i][j]+ws[i][j]-1)
		}
	}

	_, _ = fmt.Fprintf(writer, "%d", ans)
}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

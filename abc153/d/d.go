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

func answer(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	buf := make([]byte, BufferSize)
	sc.Buffer(buf, 1e+6)
	sc.Split(bufio.ScanWords)
	h := nextUint64()
	count := 1
	ans := 0
	for h > 0 {
		h = h / 2
		ans += count
		count *= 2
	}

	_, _ = fmt.Fprintf(writer, "%d", ans)
}

func main() {
	answer(os.Stdin, os.Stdout)
}

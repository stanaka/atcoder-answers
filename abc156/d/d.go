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

func mpow(x, n, M uint64) uint64 {
	ans := uint64(1)
	for n != 0 {
		if n&1 == 1 {
			ans = ans * x % M
		}
		x = x * x % M
		n = n >> 1
	}
	return ans
}

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

func answer(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	buf := make([]byte, bufferSize)
	sc.Buffer(buf, 1e+7)
	sc.Split(bufio.ScanWords)
	n := nextInt()
	a := nextInt()
	b := nextInt()

	mod := int(1e9) + 7
	ans := int(mpow(uint64(2), uint64(n), uint64(mod))) - 1

	debugf("%v\n", ans)
	nume := 1
	deno := 1
	for i := 0; i < min(a, b); i++ {
		nume *= n - i
		nume %= mod
		deno *= (i + 1)
		deno %= mod
	}
	ans -= int(uint64(nume) * mpow(uint64(deno), uint64(mod-2), uint64(mod)) % uint64(mod))
	debugf("%v\n", ans)
	if ans < 0 {
		ans += mod
	}
	for i := min(a, b); i < max(a, b); i++ {
		nume *= n - i
		nume %= mod
		deno *= (i + 1)
		deno %= mod
	}
	ans -= int(uint64(nume) * mpow(uint64(deno), uint64(mod-2), uint64(mod)) % uint64(mod))
	debugf("%v\n", ans)
	if ans < 0 {
		ans += mod
	}

	_, _ = fmt.Fprintf(writer, "%d", ans)
}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

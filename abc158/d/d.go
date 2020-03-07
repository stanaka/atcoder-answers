package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

const bufferSize = 1024

func nextString() string {
	if !sc.Scan() {
		panic(nil)
	}
	return sc.Text()
}

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

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

var production bool

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
	s := nextString()
	q := nextInt()
	rev := 0

	var pre bytes.Buffer
	var post bytes.Buffer
	for i := 0; i < q; i++ {
		a := nextInt()
		if a == 1 {
			rev++
		} else {
			f := nextInt()
			c := nextString()
			if (f == 1 && rev%2 == 0) ||
				(f == 2 && rev%2 == 1) {
				pre.WriteString(fmt.Sprintf("%s", c))
				//pre += c
			} else {
				post.WriteString(fmt.Sprintf("%s", c))
				//post += c
			}
		}
	}

	if rev%2 == 0 {
		_, _ = fmt.Fprintf(writer, "%s", reverse(pre.String()))
		_, _ = fmt.Fprintf(writer, "%s", s)
		_, _ = fmt.Fprintf(writer, "%s", post.String())
	} else {
		_, _ = fmt.Fprintf(writer, "%s", reverse(post.String()))
		_, _ = fmt.Fprintf(writer, "%s", reverse(s))
		_, _ = fmt.Fprintf(writer, "%s", pre.String())
	}
}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

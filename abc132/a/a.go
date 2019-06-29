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

var production bool

func debugf(format string, v ...interface{}) {
	if !production {
		fmt.Printf(format, v...)
	}
}

func answer(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	buf := make([]byte, BufferSize)
	sc.Buffer(buf, 1e+6)
	sc.Split(bufio.ScanWords)
	s := nextString()
	m := map[string]int{}

	for i := 0; i < len(s); i++ {
		m[string(s[i])]++
	}

	flag := true
	if len(m) != 2 {
		flag = false
	}
	for k := range m {
		if m[k] != 2 {
			flag = false
		}
	}

	if flag == true {
		_, _ = fmt.Fprint(writer, "Yes")
	} else {
		_, _ = fmt.Fprint(writer, "No")
	}
}

func main() {
	production = true
	answer(os.Stdin, os.Stdout)
}

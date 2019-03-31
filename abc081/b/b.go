package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc *bufio.Scanner

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
func answer(reader io.Reader, writer io.Writer) {
	sc = bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)
	a := nextInt()
	b := make([]int, a+1)
	for i := 0; i < a; i++ {
		b[i] = nextInt()
	}
	c := 0
	f := true
	for f == true {
		for i := 0; i < a; i++ {
			if b[i]%2 == 0 {
				b[i] = b[i] / 2
			} else {
				f = false
				break
			}
		}
		if f == true {
			c++
		}
	}

	fmt.Fprintf(writer, "%d", c)
}

func main() {
	answer(os.Stdin, os.Stdout)
}

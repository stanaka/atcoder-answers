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
	b := nextInt()
	if a%2 == 0 || b%2 == 0 {
		fmt.Fprintf(writer, "Even")
	} else {
		fmt.Fprintf(writer, "Odd")
	}
}

func main() {
	answer(os.Stdin, os.Stdout)
}

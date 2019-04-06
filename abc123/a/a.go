package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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
	a := make([]int, 5)
	for i := 0; i < 5; i++ {
		a[i] = nextInt()
	}
	k := nextInt()
	sort.Ints(a)
	if a[4]-a[0] > k {
		fmt.Fprintf(writer, ":(")
	} else {
		fmt.Fprintf(writer, "Yay!")
	}
}

func main() {
	answer(os.Stdin, os.Stdout)
}

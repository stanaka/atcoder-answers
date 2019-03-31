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
	x := -1
	y := -1
	z := -1
	c := b - a*1000
	for i := 0; i <= a && c-i*9000 >= 0; i++ {
		t := float64(c - i*9000)
		if (t/4000) == float64(int(t/4000)) && a-i-int(t/4000) >= 0 {
			x = i
			y = int(t / 4000)
			z = a - x - y
			break
		}
	}
	/*
		    x+y+z = a
			x*10000 + y*5000 + z*1000 = b
		    z = a - x - y
		    x*10000 + y*5000 + (a-x-y)*1000 = b
		    x*9000 + y*4000 = b - a*1000
		    y = (b-a*1000-x*9000)/4000
	*/

	fmt.Fprintf(writer, "%d %d %d", x, y, z)
}

func main() {
	answer(os.Stdin, os.Stdout)
}

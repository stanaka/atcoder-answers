package main

import (
	"fmt"
	"io"
	"os"
)

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func answer(reader io.Reader, writer io.Writer) {
	var n int
	_, _ = fmt.Fscan(reader, &n)
	a := make([]int, n)
	a1 := make([]int, n-1)
	a2 := make([]int, n-1)
	for i := 0; i < n; i++ {
		var v int
		_, _ = fmt.Fscan(reader, &v)
		a[i] = v
	}

	if n == 2 {
		if a[0] < a[1] {
			_, _ = fmt.Fprintf(writer, "%d", a[1])
		} else {
			_, _ = fmt.Fprintf(writer, "%d", a[0])
		}
		return
	}
	a1[0] = a[0]
	a2[n-2] = a[n-1]
	for i := 1; i < n-1; i++ {
		a1[i] = gcd(a1[i-1], a[i])
		a2[n-2-i] = gcd(a2[n-1-i], a[n-1-i])
	}
	//fmt.Printf("%v\n", a)
	//fmt.Printf("%v\n", a1)
	//fmt.Printf("%v\n", a2)
	m := a2[0]
	for i := 0; i < n-1; i++ {
		var c int
		if i < n-2 {
			c = gcd(a1[i], a2[i+1])
		} else {
			c = a1[i]
		}
		if m < c {
			m = c
		}
	}
	_, _ = fmt.Fprintf(writer, "%d", m)
}

func main() {
	answer(os.Stdin, os.Stdout)
}

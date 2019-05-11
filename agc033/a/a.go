package main

import (
	"fmt"
	"io"
	"os"
)

func answer(reader io.Reader, writer io.Writer) {
	var h, w int
	_, _ = fmt.Fscan(reader, &h)
	_, _ = fmt.Fscan(reader, &w)
	a := make([][]int, h)
	var b = map[int]map[int]bool{}
	for i := 0; i < h; i++ {
		var s string
		_, _ = fmt.Fscan(reader, &s)
		a[i] = make([]int, w)
		b[i] = map[int]bool{}
		for j := 0; j < w; j++ {
			if string(s[j]) == "#" {
				a[i][j] = 1
				b[i][j] = true
			} else {
				a[i][j] = 0
			}
		}
	}
	c := 1
	for true {
		f := false
		for i := range b {
			for j := range b[i] {
				if a[i][j] == c+1 {
					continue
				}
				delete(b[i], j)
				if i > 0 && a[i-1][j] == 0 {
					a[i-1][j] = c + 1
					b[i-1][j] = true
					f = true
				}
				if j > 0 && a[i][j-1] == 0 {
					a[i][j-1] = c + 1
					b[i][j-1] = true
					f = true
				}
				if i < h-1 && a[i+1][j] == 0 {
					a[i+1][j] = c + 1
					b[i+1][j] = true
					f = true
				}
				if j < w-1 && a[i][j+1] == 0 {
					a[i][j+1] = c + 1
					b[i][j+1] = true
					f = true
				}
			}
		}
		if f == false {
			break
		}
		c++
	}
	_, _ = fmt.Fprintf(writer, "%d", c-1)
}

func main() {
	answer(os.Stdin, os.Stdout)
}

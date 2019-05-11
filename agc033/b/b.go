package main

import (
	"fmt"
	"io"
	"os"
)

func answer(reader io.Reader, writer io.Writer) {
	var h, w, n, x, y int
	_, _ = fmt.Fscan(reader, &h)
	_, _ = fmt.Fscan(reader, &w)
	_, _ = fmt.Fscan(reader, &n)
	_, _ = fmt.Fscan(reader, &y)
	_, _ = fmt.Fscan(reader, &x)
	var s, t string
	_, _ = fmt.Fscan(reader, &s)
	_, _ = fmt.Fscan(reader, &t)

	xA := x
	xB := x
	yA := y
	yB := y
	f := true
	//fmt.Printf("x:%v, y:%v, w:%v, h:%v\n", x, y, w, h)

	for i := 0; i < n; i++ {
		if string(s[i]) == "L" {
			xA--
		}
		if string(s[i]) == "R" {
			xB++
		}
		if string(s[i]) == "U" {
			yA--
		}
		if string(s[i]) == "D" {
			yB++
		}
		if xA <= 0 || w < xB || yA <= 0 || h < yB {
			f = false
			break
		}
		if string(t[i]) == "R" && xA < w {
			xA++
		}
		if string(t[i]) == "L" && 1 < xB {
			xB--
		}
		if string(t[i]) == "D" && yA < h {
			yA++
		}
		if string(t[i]) == "U" && 1 < yB {
			yB--
		}
		//fmt.Printf("s:%v, t:%v, xA:%v, xB:%v, yA:%v, yB:%v\n", string(s[i]), string(t[i]), xA, xB, yA, yB)
		if xA <= 0 || w < xB || yA <= 0 || h < yB {
			f = false
			break
		}
	}

	if f == true {
		_, _ = fmt.Fprintf(writer, "YES")
	} else {
		_, _ = fmt.Fprintf(writer, "NO")
	}
}

func main() {
	answer(os.Stdin, os.Stdout)
}

package main

import (
	"fmt"
	"io"
	"os"
)

func answer(reader io.Reader, writer io.Writer) {
	var x, y, r, g, b, n int
	_, _ = fmt.Fscan(reader, &r)
	_, _ = fmt.Fscan(reader, &g)
	_, _ = fmt.Fscan(reader, &b)
	_, _ = fmt.Fscan(reader, &n)
	count := 0
	x = n / r
	for i:=x;i>=0;i-- {
		if i*r == n {
			count++
			//_, _ = fmt.Printf("%d: %d, 0, 0\n", count, i)
		} else {
			y = (n-i*r) / g
			for j:=y;j>=0;j-- {
				if i*r + j*g == n {
					count++
					//_, _ = fmt.Printf("%d: %d, %d, 0\n", count, i, j)
				} else {
					if (n-i*r-j*g) % b == 0 {
						count++
						//_, _ = fmt.Printf("%d: %d, %d, %d\n", count, i, j,(n-i*r-j*g)/b)
					}
				}
			}

		}
	}

	//_, _ = fmt.Printf("%d\n", count)
	_, _ = fmt.Fprintf(writer, "%d", count)
}

func main() {
	answer(os.Stdin, os.Stdout)
}

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
	count1 :=0
	count2 :=0
	count3 :=0
	h :=0
	t :=0
	s := make([]string, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscan(reader, &s[i])
		f :=0
		if s[i][0] == 'B' {
			h++
			f=1
		}
		if s[i][len(s[i])-1] == 'A' {
			t++
			f=1
		}
		if f==1 {
			count1++
		}
	}
	for i := 0; i < n; i++ {
		for j:=0;j < len(s[i])-1;j++ {
			if s[i][j] == 'A' {
				if s[i][j+1] == 'B' {
					count2++
				}
			}
		}
	}
	if t < h {
		count3 = t
	} else {
		count3 = h
	}
	if count3 > 0 && count1 == count3 {
		count3--
	}

	_, _ = fmt.Fprintf(writer, "%d", count2 + count3)
}

func main() {
	answer(os.Stdin, os.Stdout)
}

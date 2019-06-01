package main

import (
	"fmt"
	"io"
	"os"
)

func answer(reader io.Reader, writer io.Writer) {
	var s string
	_, _ = fmt.Fscan(reader, &s)
	c:=0
	for i:=0; i<len(s); i++ {
		if string(s[i]) == "x" {
			c++
		}
	}
	if c <= 7 {
		_, _ = fmt.Fprint(writer,"YES")
	} else {
		_, _ = fmt.Fprint(writer,"NO")
	}
}

func main() {
	answer(os.Stdin, os.Stdout)
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func answer(reader io.Reader, writer io.Writer) {
	r := bufio.NewReader(reader)
	a := 0
	n := "0"
	b := 0
	var s string
	for {
		if c, _, err := r.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			s = string(c)
			if s != "0" && s != "1" {
				break
			}
		}
		if s == n {
			b++
		} else {
			a++
		}
		if n == "0" {
			n = "1"
		} else {
			n = "0"
		}
	}
	if b < a {
		a = b
	}
	_, _ = fmt.Fprintf(writer, "%d\n", a)
}

func main() {
	answer(os.Stdin, os.Stdout)
}

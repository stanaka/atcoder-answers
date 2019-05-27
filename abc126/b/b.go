package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func answer(reader io.Reader, writer io.Writer) {
	var s string
	_, _ = fmt.Fscan(reader, &s)
	var a, b int
	a,_ = strconv.Atoi(string(s[:2]))
	b,_ = strconv.Atoi(s[2:])
	var yymm, mmyy bool
	yymm = false
	mmyy = false
	//fmt.Printf("s:%s, a:%d, b:%d\n",s,a,b)
	if 0 < a && a < 13 {
		// a might be month
		mmyy = true
	}
	if 0 < b && b < 13 {
		// a might be month
		yymm = true
	}
	var r string
	r = "NA"
	if mmyy == true && yymm == true {
		r = "AMBIGUOUS"
	} else if mmyy == true && yymm == false {
		r = "MMYY"
	} else if mmyy == false && yymm == true {
		r = "YYMM"
	}
	_, _ = fmt.Fprintf(writer, "%s", r)
}

func main() {
	answer(os.Stdin, os.Stdout)
}

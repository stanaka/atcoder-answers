package main

import (
	"io"
	"sort"
)

type bc struct {
	num   int
	value int
}

type bcs []bc

func (b bcs) Len() int           { return len(b) }
func (b bcs) Less(i, j int) bool { return b[i].value < b[j].value }
func (b bcs) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

func answer(reader io.Reader, writer io.Writer) {
	sort.Sort(sort.Reverse(bcs(b)))
}

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

	sort.Ints(p)
	sort.Sort(sort.Reverse(sort.IntSlice(m)))

	// binary search for sorted array
	j := sort.Search(len(a), func(k int) bool { return a[k] > x })
	// binary search for reversed sorted array
	j := sort.Search(len(a), func(k int) bool { return a[k] < x })
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

type sortStrings []string

func (s sortStrings) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortStrings) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortStrings) Len() int {
	return len(s)
}

func SortStrings(s []string) []string {
	sort.Sort(sortStrings(s))
	return s
}

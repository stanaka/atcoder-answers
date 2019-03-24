package utils

import (
	"bytes"
	"io"
	"strings"
)

func Helper(input string, f func(io.Reader, io.Writer)) string {
	r := strings.NewReader(input)
	w := bytes.NewBufferString("")
	f(r, w)
	return w.String()
}

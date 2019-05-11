package main

import (
	"testing"

	"github.com/stanaka/atcoder-answers/utils"
)

func TestAnswer1(t *testing.T) {
	input := `3 3
...
.#.
...`
	expect := `2`

	output := utils.Helper(input, answer)
	if output != expect {
		t.Errorf("expect %v, but %v", expect, output)
	}
}

func TestAnswer2(t *testing.T) {
	input := `6 6
..#..#
......
#..#..
......
.#....
....#.`
	expect := `3`

	output := utils.Helper(input, answer)
	if output != expect {
		t.Errorf("expect %v, but %v", expect, output)
	}
}

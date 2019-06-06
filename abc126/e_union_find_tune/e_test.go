package main

import (
	"testing"

	"github.com/stanaka/atcoder-answers/utils"
)

func TestAnswer1(t *testing.T) {
	input := `3 1
1 2 1`
	expect := `2`

	output := utils.Helper(input, answer)
	if output != expect {
		t.Errorf("expect %v, but %v", expect, output)
	}
}

func TestAnswer2(t *testing.T) {
	input := `6 5
1 2 1
2 3 2
1 3 3
4 5 4
5 6 5`
	expect := `2`

	output := utils.Helper(input, answer)
	if output != expect {
		t.Errorf("expect %v, but %v", expect, output)
	}
}

func TestAnswer2_2(t *testing.T) {
	input := `6 4
2 1 1
3 1 3
6 4 4
5 6 5`
	expect := `2`

	output := utils.Helper(input, answer)
	if output != expect {
		t.Errorf("expect %v, but %v", expect, output)
	}
}

func TestAnswer2_3(t *testing.T) {
	input := `6 3
2 1 1
3 1 3
5 6 5`
	expect := `3`

	output := utils.Helper(input, answer)
	if output != expect {
		t.Errorf("expect %v, but %v", expect, output)
	}
}

/*
func TestAnswer3(t *testing.T) {
	input := `100000 1
1 100000 100`
	expect := `99999`

	output := utils.Helper(input, answer)
	if output != expect {
		t.Errorf("expect %v, but %v", expect, output)
	}
}
*/

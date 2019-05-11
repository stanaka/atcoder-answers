package main

import (
	"testing"

	"github.com/stanaka/atcoder-answers/utils"
)

func TestAnswer1(t *testing.T) {
	input := `2 3 3
2 2
RRL
LUD`
	expect := `YES`

	output := utils.Helper(input, answer)
	if output != expect {
		t.Errorf("expect %v, but %v", expect, output)
	}
}

func TestAnswer2(t *testing.T) {
	input := `4 3 5
2 2
UDRRR
LLDUD`
	expect := `NO`

	output := utils.Helper(input, answer)
	if output != expect {
		t.Errorf("expect %v, but %v", expect, output)
	}
}

func TestAnswer3(t *testing.T) {
	input := `5 6 11
2 1
RLDRRUDDLRL
URRDRLLDLRD`
	expect := `NO`

	output := utils.Helper(input, answer)
	if output != expect {
		t.Errorf("expect %v, but %v", expect, output)
	}
}

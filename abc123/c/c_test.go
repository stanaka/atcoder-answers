package main

import (
	"testing"

	"github.com/stanaka/atcoder-answers/utils"
)

func TestAnswer1(t *testing.T) {
	input := `5
3
2
4
3
5`
	expect := `7`

	output := utils.Helper(input, answer)
	if output != expect {
		t.Errorf("expect %v, but %v", expect, output)
	}
}

func TestAnswer2(t *testing.T) {
	input := `10
123
123
123
123
123`
	expect := `5`

	output := utils.Helper(input, answer)
	if output != expect {
		t.Errorf("expect %v, but %v", expect, output)
	}
}

func TestAnswer3(t *testing.T) {
	input := `10000000007
2
3
5
7
11`
	expect := `5000000008`

	output := utils.Helper(input, answer)
	if output != expect {
		t.Errorf("expect %v, but %v", expect, output)
	}
}

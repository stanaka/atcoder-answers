package main

import (
	"testing"

	"github.com/stanaka/atcoder-answers/utils"
)

func TestAnswer1(t *testing.T) {
	input := `5
1 2
2 3
3 4
4 5
1 2 3 4 5`
	/*	expect := `10
		1 2 3 4 5` */
	expect := `10
5 4 3 2 1`

	output := utils.Helper(input, answer)
	if output != expect {
		t.Errorf("expect %v, but %v", expect, output)
	}
}

func TestAnswer2(t *testing.T) {
	input := `5
1 2
1 3
1 4
1 5
3141 59 26 53 59`
	/*	expect := `197
		59 26 3141 59 53`*/
	expect := `197
3141 59 59 53 26`

	output := utils.Helper(input, answer)
	if output != expect {
		t.Errorf("expect %v, but %v", expect, output)
	}
}

func TestAnswer4(t *testing.T) {
	input := `9
1 4
1 3
1 6
1 9
3 2
6 8
2 7
2 5
5 6 7 8 2 4 3 9 1`
	expect := `36
9 4 7 8 1 6 2 3 5`

	output := utils.Helper(input, answer)
	if output != expect {
		t.Errorf("expect %v, but %v", expect, output)
	}
}

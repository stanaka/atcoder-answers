package main

import (
	"testing"

	"github.com/stanaka/atcoder-answers/utils"
)

func TestAnswer1(t *testing.T) {
	input := `4 6
1 3 2
7 13 10
18 20 13
3 4 2
0
1
2
3
5
8`
	expect := `2
2
10
-1
13
-1
`

	output := utils.Helper(input, answer)
	if output != expect {
		t.Errorf("expect %v, but %v", expect, output)
	}
}

func TestAnswer2(t *testing.T) {
	input := `2 6
5 6 2
12 16 10
0
1
2
3
5
8`
	expect := `-1
-1
10
2
10
-1
`

	output := utils.Helper(input, answer)
	if output != expect {
		t.Errorf("expect %v, but %v", expect, output)
	}
}

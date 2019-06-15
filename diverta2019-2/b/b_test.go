package main

import (
	"testing"

	"github.com/stanaka/atcoder-answers/utils"
)

func TestAnswer1(t *testing.T) {
	input := `2
1 1
2 2`
	expect := `1`

	output := utils.Helper(input, answer)
	if output != expect {
		t.Errorf("expect %v, but %v", expect, output)
	}
}

func TestAnswer2(t *testing.T) {
	input := `3
1 4
4 6
7 8`
	expect := `1`

	output := utils.Helper(input, answer)
	if output != expect {
		t.Errorf("expect %v, but %v", expect, output)
	}
}

func TestAnswer2_2(t *testing.T) {
	input := `3
1 4
7 8
4 6`
	expect := `1`

	output := utils.Helper(input, answer)
	if output != expect {
		t.Errorf("expect %v, but %v", expect, output)
	}
}

func TestAnswer2_3(t *testing.T) {
	input := `3
1 4
7 4
4 6`
	expect := `2`

	output := utils.Helper(input, answer)
	if output != expect {
		t.Errorf("expect %v, but %v", expect, output)
	}
}

func TestAnswer2_4(t *testing.T) {
	input := `1
1 1`
	expect := `1`

	output := utils.Helper(input, answer)
	if output != expect {
		t.Errorf("expect %v, but %v", expect, output)
	}
}

func TestAnswer2_5(t *testing.T) {
	input := `3
1 4
1 0
1 2`
	expect := `1`

	output := utils.Helper(input, answer)
	if output != expect {
		t.Errorf("expect %v, but %v", expect, output)
	}
}

func TestAnswer2_6(t *testing.T) {
	input := `6
1 4
1 0
1 2
4 1
0 1
2 1`
	expect := `4`

	output := utils.Helper(input, answer)
	if output != expect {
		t.Errorf("expect %v, but %v", expect, output)
	}
}

func TestAnswer3(t *testing.T) {
	input := `4
1 1
1 2
2 1
2 2`
	expect := `2`

	output := utils.Helper(input, answer)
	if output != expect {
		t.Errorf("expect %v, but %v", expect, output)
	}
}

/*
func TestAnswerGenerate1(t *testing.T) {
	input := `100000
`
	for i := 0; i < 100000; i++ {
		input += "0 "
	}
	expect := `0`

	s := time.Now().UnixNano()
	fmt.Printf("start:%d\n", time.Now().UnixNano()/1000000)
	output := utils.Helper(input, answer)
	fmt.Printf("end  :%d (%d msec)\n", time.Now().UnixNano()/1000000, (time.Now().UnixNano()-s)/1000000)
	if output != expect {
		t.Errorf("expect %v, but %v", expect, output)
	}
}
*/

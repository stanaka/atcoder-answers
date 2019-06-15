package main

import (
	"testing"

	"github.com/stanaka/atcoder-answers/utils"
)

func TestAnswer1(t *testing.T) {
	input := `3
1 -1 2`
	expect := `4
-1 1
2 -2
`

	output := utils.Helper(input, answer)
	if output != expect {
		t.Errorf("expect %v, but %v", expect, output)
	}
}

func TestAnswer2(t *testing.T) {
	input := `3
1 1 1`
	expect := `1
1 1
1 0
`

	output := utils.Helper(input, answer)
	if output != expect {
		t.Errorf("expect %v, but %v", expect, output)
	}
}

func TestAnswer3(t *testing.T) {
	input := `3
-1 -2 -4`
	expect := `5
-1 -4
3 -2
`

	output := utils.Helper(input, answer)
	if output != expect {
		t.Errorf("expect %v, but %v", expect, output)
	}
}

func TestAnswer4(t *testing.T) {
	input := `3
-1 -2 4`
	expect := `7
4 -2
6 -1
`

	output := utils.Helper(input, answer)
	if output != expect {
		t.Errorf("expect %v, but %v", expect, output)
	}
}

func TestAnswer5(t *testing.T) {
	input := `4
1 4 5 6`
	expect := `14
1 4
-3 5
6 -8
`

	output := utils.Helper(input, answer)
	if output != expect {
		t.Errorf("expect %v, but %v", expect, output)
	}
}

func TestAnswer6(t *testing.T) {
	input := `5
-1 0 4 5 6`
	expect := `16
-1 0
-1 4
-5 5
6 -10
`

	output := utils.Helper(input, answer)
	if output != expect {
		t.Errorf("expect %v, but %v", expect, output)
	}
}

func TestAnswer7(t *testing.T) {
	input := `3
0 -2 -4`
	expect := `6
0 -4
4 -2
`

	output := utils.Helper(input, answer)
	if output != expect {
		t.Errorf("expect %v, but %v", expect, output)
	}
}

func TestAnswer8(t *testing.T) {
	input := `3
0 0 0`
	expect := `0
0 0
0 0
`

	output := utils.Helper(input, answer)
	if output != expect {
		t.Errorf("expect %v, but %v", expect, output)
	}
}

func TestAnswer9(t *testing.T) {
	input := `3
0 0 -1`
	expect := `1
-1 0
0 -1
`

	output := utils.Helper(input, answer)
	if output != expect {
		t.Errorf("expect %v, but %v", expect, output)
	}
}

func TestAnswer10(t *testing.T) {
	input := `4
-1 -4 -5 -6`
	expect := `14
-1 -6
5 -5
10 -4
`

	output := utils.Helper(input, answer)
	if output != expect {
		t.Errorf("expect %v, but %v", expect, output)
	}
}

func TestAnswer11(t *testing.T) {
	input := `2
0 2`
	expect := `2
2 0
`

	output := utils.Helper(input, answer)
	if output != expect {
		t.Errorf("expect %v, but %v", expect, output)
	}
}

func TestAnswer12(t *testing.T) {
	input := `2
-1 -2`
	expect := `1
-1 -2
`

	output := utils.Helper(input, answer)
	if output != expect {
		t.Errorf("expect %v, but %v", expect, output)
	}
}

// 124 -> -14 -> 5
// 1247 -> -147 -> -57 -> 12
// 1247 -> -327 -> -57 -> 12
// 1247 -> -624 -> -84 -> 12
// 1247 -> -312 -> -42 -> -6xx
// 1-247 -> -347 -> -77-> 14
// 1-247 -> -914 -> -10,4 -> -14
// 1-247 -> 1-67 -> -77 -> 14
// 1-247 -> -914 -> -10,4 -> 14
// -1-2-4 -> 1-4 -> 5
// -1-24 -> -16 -> 7
// 1456 -> -356 -> -86 -> 14
// 1456 -> -446 -> -86 -> 14

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

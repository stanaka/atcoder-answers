package main

import (
	"github.com/stanaka/atcoder-answers/utils"
	"testing"
)

type testCase struct {
	input  string
	expect string
}

func TestAnswer(t *testing.T) {
	testCases := []testCase{
		{`2 2
1 3
3 1`, `3`},
		{`2 2
1 1
1 1`, `6`},
		{`4 4
3 4 5 6
3 4 5 6`, `16`},
		{`10 9
9 6 5 7 5 9 8 5 6 7
8 6 8 5 5 7 9 9 7`, `191`},
		{`20 20
1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1
1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1`, `846527861`},
	}

	for i, v := range testCases {
		output := utils.Helper(v.input, answer)
		if output != v.expect {
			t.Errorf("case[%d]: expect %v, but %v", i+1, v.expect, output)
		}
	}
}

/*
func TestAnswerGenerate1(t *testing.T) {
	a := 400
	input := fmt.Sprintf("%d %d\n", a, a)
	for i := 0; i < a; i++ {
		input += fmt.Sprintf("%d ", i)
	}
	input += "\n"
	for i := 0; i < a; i++ {
		input += fmt.Sprintf("%d ", i)
	}
	expect := `0`
	if a == 400 {
		expect = `198967538`
	}

	s := time.Now().UnixNano()
	fmt.Printf("start:%d\n", time.Now().UnixNano()/1000000)
	output := utils.Helper(input, answer)
	fmt.Printf("end  :%d (%d msec)\n", time.Now().UnixNano()/1000000, (time.Now().UnixNano()-s)/1000000)
	if output != expect {
		t.Errorf("expect %v, but %v", expect, output)
	}
}
*/

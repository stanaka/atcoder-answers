package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stanaka/atcoder-answers/utils"
)

type testCase struct {
	input  string
	expect string
}

func TestAnswer(t *testing.T) {
	testCases := []testCase{
		{`contest
son`, `10`},
		{`contest
programming`, `-1`},
		{`contest
sentence`, `33`},
		{`contest
sentetnce`, `33`},
		{`contest
ne`, `5`},
		{`ss
sss`, `3`},
	}

	for i, v := range testCases {
		output := utils.Helper(v.input, answer)
		if output != v.expect {
			t.Errorf("case[%d]: expect %v, but %v", i+1, v.expect, output)
		}
	}
}

func TestAnswerGenerate1(t *testing.T) {
	var input string
	for i := 0; i < 100000; i++ {
		input += "s"
	}
	input += fmt.Sprint("\n")
	for i := 0; i < 100000; i++ {
		input += "s"
	}
	expect := `5`

	s := time.Now().UnixNano()
	fmt.Printf("start:%d\n", time.Now().UnixNano()/1000000)
	production = true
	output := utils.Helper(input, answer)
	//output := expect
	fmt.Printf("end  :%d (%d msec)\n", time.Now().UnixNano()/1000000, (time.Now().UnixNano()-s)/1000000)
	if output != expect {
		t.Errorf("expect %v, but %v", expect, output)
	}
}

package main

import (
	"testing"

	"github.com/stanaka/atcoder-answers/utils"
)

type testCase struct {
	input  string
	expect string
}

func TestAnswer(t *testing.T) {
	testCases := []testCase{
		{`4 3
1 2
2 3
2 4
2 10
1 100
3 1`, `100 110 111 110`},
		{`6 2
1 2
1 3
2 4
3 6
2 5
1 10
1 10`, `20 20 20 20 20 20`},
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

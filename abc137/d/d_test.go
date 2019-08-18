package main

import (
	"fmt"
	"github.com/stanaka/atcoder-answers/utils"
	"testing"
	"time"
)

type testCase struct {
	input  string
	expect string
}

func TestAnswer(t *testing.T) {
	testCases := []testCase{
		{`3 4
4 3
4 1
2 2`, `5`},
		{`5 3
1 2
1 3
1 4
2 1
2 3`, `10`},
		{`1 1
2 1`, `0`},
		{`2 2
2 1
1 3`, `4`},
	}

	for i, v := range testCases {
		output := utils.Helper(v.input, answer)
		if output != v.expect {
			t.Errorf("case[%d]: expect %v, but %v", i+1, v.expect, output)
		}
	}
}

func TestAnswerGenerate1(t *testing.T) {
	num := int(1e5)
	input := ""
	input += fmt.Sprintf("%d %d\n", num, num)
	for i := 0; i < num; i++ {
		input += fmt.Sprintf("%d %d\n", 1, 1)
	}
	expect := `0`

	s := time.Now().UnixNano()
	fmt.Printf("start:%d\n", time.Now().UnixNano()/1000000)
	output := utils.Helper(input, answer)
	fmt.Printf("end  :%d (%d msec)\n", time.Now().UnixNano()/1000000, (time.Now().UnixNano()-s)/1000000)
	if output != expect {
		//t.Errorf("expect %v, but %v", expect, output)
	}
}

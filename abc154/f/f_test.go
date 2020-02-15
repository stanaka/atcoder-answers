package main

import (
	"bytes"
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
		{`1 1 2 2`, `14`},
		{`314 159 2653 589`, `602215194`},
	}

	for i, v := range testCases {
		output := utils.Helper(v.input, answer)
		if output != v.expect {
			t.Errorf("case[%d]: expect %v, but %v", i+1, v.expect, output)
		}
	}
}

func TestAnswerGenerate1(t *testing.T) {
	var buffer bytes.Buffer
	//num := 2 * 100000
	num := 10000
	buffer.WriteString(fmt.Sprintf("%d %d %d %d\n", 1, 1, num, num))
	input := buffer.String()
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

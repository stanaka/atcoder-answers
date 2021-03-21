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
		{`3
2 8 4`, `56`},
		{`5
-5 8 9 -4 -3`, `950`},
		// {``, ``},
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
	num := 3 * int(1e5)
	buffer.WriteString(fmt.Sprintf("%d\n", num))
	for i := 0; i < num; i++ {
		buffer.WriteString(fmt.Sprintf("%d ", i%401-200))
	}
	buffer.WriteString("\n")
	input := buffer.String()
	expect := `1206269415462524`

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

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
		{`4 4 1
2 1
1 3
3 2
3 4
4 1`, `0 1 0 1`},
		{`5 10 0
1 2
1 3
1 4
1 5
3 2
2 4
2 5
4 3
5 3
4 5`, `0 0 0 0 0`},
		{`10 9 3
10 1
6 7
8 2
2 5
8 4
7 3
10 9
6 4
5 8
2 6
7 5
3 1`, `1 3 5 4 3 3 3 3 1 0`},
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
	num := 100000
	//buffer.WriteString(fmt.Sprintf("%d %d %d\n", num, num, num))
	buffer.WriteString(fmt.Sprintf("%d %d %d\n", num, 0, 0))
	for i := 0; i < num; i++ {
		buffer.WriteString(fmt.Sprintf("%d %d\n", i+1, ((i+1)%num)+1))
	}
	for i := 0; i < num; i++ {
		buffer.WriteString(fmt.Sprintf("%d %d\n", i+1, ((i+2)%num)+1))
	}
	input := buffer.String()
	expect := `5`

	s := time.Now().UnixNano()
	fmt.Printf("start:%d (length: %d)\n", time.Now().UnixNano()/1000000, len(input))
	production = true
	output := utils.Helper(input, answer)
	//output := expect
	fmt.Printf("end  :%d (%d msec)\n", time.Now().UnixNano()/1000000, (time.Now().UnixNano()-s)/1000000)
	if output != expect {
		//t.Errorf("expect %v, but %v", expect, output)
	}
}

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
		{`3 1
3 1 4
5 3 2`, `1`},
		{`7 3
27 18 28 18 28 46 1000000000
1000000000 1 7
1000000000 2 10
1000000000 3 12`, `224489796
214285714
559523809`},
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
	var buffer bytes.Buffer
	//num := 2 * 100000
	num := 2 * 100000
	buffer.WriteString(fmt.Sprintf("%d %d %d\n", num, num, 1))
	for i := 0; i < num; i++ {
		buffer.WriteString(fmt.Sprintf("%d %d\n", i, i+1))
	}
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
*/

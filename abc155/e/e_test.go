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
		{`36`, `8`},
		{`91`, `3`},
		{`5`, `5`},
		{`56`, `9`},    // 100->044, 60->4
		{`455`, `14`},  // 500->045
		{`55`, `10`},   // 100->045, 60->5
		{`65`, `9`},    // 100->035, 105->40
		{`45`, `9`},    // 100->055, 45->00
		{`54`, `9`},    // 54->00
		{`3755`, `15`}, // 4000->0245, 4055->0300
		{`314159265358979323846264338327950288419716939937551058209749445923078164062862089986280348253421170`, `243`},
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

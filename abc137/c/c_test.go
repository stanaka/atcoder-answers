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
		{`3
acornistnt
peanutbomb
constraint`, `1`},
		{`2
oneplustwo
ninemodsix`, `0`},
		{`5
abaaaaaaaa
oneplustwo
aaaaaaaaba
twoplusone
aaaabaaaaa`, `4`},
		{`5
acaaaaaaaa
abaaaaaaaa
aaaaaaaaaa
aaaaaaaaba
aaaabaaaaa`, `3`},
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

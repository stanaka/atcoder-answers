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
		{`3 4
		5 1
		10 1
		8 0
		1 10
		4 5
		6 7
		8 9`, `2
1 4`},
		{`4 2
		2 0
		3 1
		5 1
		7 0
		1 4
		4 7`, `-1`},
		{`3 2
		5 0
		10 0
		8 0
		6 9
		66 99`, `0
`},
		{`12 20
		536130100 1
		150049660 1
		79245447 1
		132551741 0
		89484841 1
		328129089 0
		623467741 0
		248785745 0
		421631475 0
		498966877 0
		43768791 1
		112237273 0
		21499042 142460201
		58176487 384985131
		88563042 144788076
		120198276 497115965
		134867387 563350571
		211946499 458996604
		233934566 297258009
		335674184 555985828
		414601661 520203502
		101135608 501051309
		90972258 300372385
		255474956 630621190
		436210625 517850028
		145652401 192476406
		377607297 520655694
		244404406 304034433
		112237273 359737255
		392593015 463983307
		150586788 504362212
		54772353 83124235`, `5
1 7 8 9 11`},
	}

	for i, v := range testCases {
		output := utils.Helper(v.input, answer)
		if output != v.expect {
			t.Errorf("case[%d]: expect '%v', but '%v'", i+1, v.expect, output)
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

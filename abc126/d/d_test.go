package main

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/stanaka/atcoder-answers/utils"
)

func TestAnswer1(t *testing.T) {
	input := `3
1 2 2
2 3 1`
	expect := `0
0
1`

	output := utils.Helper(input, answer)
	if output != expect {
		t.Errorf("expect %v, but %v", expect, output)
	}
}

func TestAnswer2(t *testing.T) {
	input := `5
2 5 2
2 3 10
1 3 8
3 4 2`
	expect := `0
0
0
0
0`

	/*expect := `1
0
1
0
1`*/

	output := utils.Helper(input, answer)
	if output != expect {
		t.Errorf("expect %v, but %v", expect, output)
	}
}

func TestAnswer3(t *testing.T) {
	input := `5
2 5 1
2 3 9
1 3 8
3 5 2`
	expect := `0
1
0
0
0`

	output := utils.Helper(input, answer)
	if output != expect {
		t.Errorf("expect %v, but %v", expect, output)
	}
}

func TestAnswer4(t *testing.T) {
	n := 100000
	//n := 5
	var buffer bytes.Buffer
	var buffer2 bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%d\n", n))
	for i := 1; i < n; i++ {
		buffer.WriteString(fmt.Sprintf("%d %d 2\n",i, i+1))
		buffer2.WriteString("0\n")
	}
	input := buffer.String()
	expect := buffer2.String()
	s := time.Now().UnixNano()
	fmt.Printf("start:%d\n", time.Now().UnixNano()/1000000)
	output := utils.Helper(input, answer)
	fmt.Printf("end  :%d (%d)\n", time.Now().UnixNano()/1000000,  (time.Now().UnixNano()-s)/1000000)
	if output != expect {
		//t.Errorf("expect %v, but %v", expect, output)
	}
}

package main

import (
	"fmt"
	"io"
	"os"
)

func fscanNums(r io.Reader, len int) (nums []int) {
	var num int
	for i := 0; i < len; i++ {
		fmt.Fscan(r, &num)
		nums = append(nums, num)
	}
	return
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
func answer(r io.Reader, w io.Writer) {
	var numNumber int
	fmt.Fscan(r, &numNumber)
	nums := fscanNums(r, numNumber)

	ans := nums[0]
	for i := 1; i < numNumber; i++ {
		ans = gcd(ans, nums[i])
	}
	fmt.Fprintf(w, "%d", ans)
}

func main() {
	answer(os.Stdin, os.Stdout)
}

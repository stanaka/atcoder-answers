package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

func fscanNums(r io.Reader, len int) (nums []int) {
	var num int
	for i := 0; i < len; i++ {
		fmt.Fscan(r, &num)
		nums = append(nums, num)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func answer(r io.Reader, w io.Writer) {

	msSizes := []int{0, 2, 5, 5, 4, 5, 6, 3, 7, 6}

	var numMS int
	var numDigits int
	fmt.Fscan(r, &numMS)
	fmt.Fscan(r, &numDigits)
	nums := fscanNums(r, numDigits)
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	dp := make([]int, numMS+1)
	dp[0] = 0
	for i := 1; i <= numMS; i++ {
		maxDigit := -1
		for j := 0; j < numDigits; j++ {
			index := i - msSizes[nums[j]]
			if index >= 0 && dp[index] >= 0 {
				maxDigit = max(maxDigit, dp[index]+1)
			}
		}
		dp[i] = maxDigit
	}

	remain := numMS
	for remain > 0 {
		maxDigit := -1
		for j := 0; j < numDigits; j++ {
			index := remain - msSizes[nums[j]]

			if index >= 0 && dp[index] == dp[remain]-1 {
				maxDigit = max(maxDigit, nums[j])
			}
		}
		remain = remain - msSizes[maxDigit]
		fmt.Fprintf(w, "%v", maxDigit)
	}
}

func main() {
	answer(os.Stdin, os.Stdout)
}

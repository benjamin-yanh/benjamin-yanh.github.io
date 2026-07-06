package top100liked_golang

import (
	"sort"
	"testing"
)

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	maxLen := 1
	size := 1
	preNum := nums[0]

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if num > preNum && num-preNum == 1 {
			size++
		} else if num == preNum {
			continue
		} else {
			size = 1
		}
		preNum = num
		if size > maxLen {
			maxLen = size
		}
	}

	return maxLen
}

func Test_longestConsecutive(t *testing.T) {

}

package top100liked_golang

import (
	"fmt"
	"testing"
)

func maxSubArray(nums []int) int {
	sum := nums[0]
	val := sum
	for i := 1; i < len(nums); i++ {
		if sum < 0 && sum < nums[i] {
			sum = nums[i]
		} else {
			sum += nums[i]
		}
		val = max(val, sum)
	}
	return val
}

func Test_maxSubArray(t *testing.T) {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{1}))
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))
	fmt.Println(maxSubArray([]int{1, 2}))
}

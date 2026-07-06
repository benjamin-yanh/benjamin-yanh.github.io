package top100liked_golang

import "testing"

func Test_L1(t *testing.T) {
	twoSum([]int{1, 2, 3}, 6)
}

func twoSum(nums []int, target int) []int {
	var answer [2]int
	maps := make(map[int]int)
	for i, num := range nums {
		maps[num] = i
	}
}

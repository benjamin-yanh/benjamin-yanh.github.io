package top100liked_golang

import (
	"fmt"
	"sort"
	"testing"
)

func Test_L1(t *testing.T) {
	twoSum([]int{1, 2, 3}, 6)
	fmt.Println(sort.Reverse(sort.IntSlice([]int{-1, -2, -3, -4, -5})))
}

func twoSum(nums []int, target int) []int {
	nums = sort.IntSlice(nums)
	maps := make(map[int]int)
	for i, num := range nums {
		if index, exist := maps[target-num]; exist {
			return []int{index, i}
		}
		maps[num] = i
	}
	return []int{0, 0}
}

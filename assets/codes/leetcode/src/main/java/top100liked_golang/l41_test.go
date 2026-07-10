package top100liked_golang

import (
	"fmt"
	"sort"
	"testing"
)

func firstMissingPositive(nums []int) int {
	sort.Ints(nums)
	maps := make(map[int]int)
	for _, num := range nums {
		if num <= 0 {
			continue
		}
		maps[num]++
		if num != len(maps) {
			return len(maps)
		}
	}

	return len(maps) + 1
}

func Test_firstMissingPositive(t *testing.T) {
	fmt.Println(firstMissingPositive([]int{7, 8, 9, 11, 12}))
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))
}

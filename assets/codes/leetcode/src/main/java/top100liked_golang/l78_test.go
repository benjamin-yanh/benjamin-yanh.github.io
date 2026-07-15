package top100liked_golang

import (
	"fmt"
	"testing"
)

func subsets(nums []int) [][]int {
	numsLen := len(nums)
	result := [][]int{}
	result = append(result, []int{})
	for i := 1; i <= numsLen; i++ {
		findSubsets(nums, []int{}, i, &result)
	}
	return result
}

func findSubsets(nums, ans []int, length int, result *[][]int) {
	if len(ans) == length {
		c := []int{}
		c = append(c, ans...)
		*result = append(*result, c)
		return
	}
	for i := 0; i < len(nums); i++ {
		ans = append(ans, nums[i])
		arr := append([]int{}, nums[i+1:]...)
		findSubsets(arr, ans, length, result)
		ans = ans[:len(ans)-1]
	}

}

func Test_subsets(t *testing.T) {
	fmt.Println(subsets([]int{1, 2, 3}))
}

package top100liked_golang

import (
	"fmt"
	"testing"
)

func permute(nums []int) [][]int {
	result := [][]int{}
	permuteAns(nums, []int{}, &result)
	return result
}

func permuteAns(nums, ans []int, result *[][]int) {
	if len(nums) <= 0 {
		value := make([]int, len(ans))
		copy(value, ans)
		*result = append(*result, value)
		return
	}
	for i := 0; i < len(nums); i++ {
		ans = append(ans, nums[i])
		arr := []int{}
		arr = append(arr, nums[:i]...)
		arr = append(arr, nums[i+1:]...)
		permuteAns(arr, ans, result)
		ans = ans[:len(ans)-1]
	}

}

func Test_permute(t *testing.T) {
	fmt.Println(permute([]int{1, 2, 3}))
}

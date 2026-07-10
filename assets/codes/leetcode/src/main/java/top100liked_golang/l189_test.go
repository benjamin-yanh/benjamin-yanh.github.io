package top100liked_golang

import (
	"fmt"
	"testing"
)

func rotate(nums []int, k int) {
	length := len(nums)
	k = k % length
	ans := make([]int, length)
	copy(ans, nums)
	vals := ans[length-k:]
	ans = append(vals, ans...)
	copy(nums, ans)

	return
}

func Test_rotate(t *testing.T) {
	//{
	//	arr := []int{1, 2, 3, 4, 5, 6, 7}
	//	rotate(arr, 3)
	//	fmt.Println(arr)
	//}
	{
		arr := []int{-1}
		rotate(arr, 2)
		fmt.Println(arr)
	}
}

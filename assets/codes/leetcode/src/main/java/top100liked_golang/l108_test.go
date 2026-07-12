package top100liked_golang

import (
	"fmt"
	"testing"
)

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	length := len(nums)
	return &TreeNode{Val: nums[length/2], Left: sortedArrayToBST(nums[:length/2]), Right: sortedArrayToBST(nums[(length/2)+1:])}
}

func Test_sortedArrayToBST(t *testing.T) {
	val := sortedArrayToBST([]int{0, 1, 2, 3, 4, 5})
	fmt.Println(val)
}

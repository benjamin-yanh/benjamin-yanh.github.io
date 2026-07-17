package top100liked_golang

import (
	"slices"
	"testing"
)

func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	if n == 1 {
		return nums[0]
	}
	left, right := 0, n-1

	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[0] < nums[mid] {
			if nums[0] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[n-1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return slices.Index(nums, target)
}

func Test_search(t *testing.T) {

}

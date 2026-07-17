package top100liked_golang

import (
	"fmt"
	"testing"
)

func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	ans := []int{-1, -1}
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			l := mid
			r := mid
			for l >= 0 && nums[l] == target {
				l--
			}
			for r < len(nums) && nums[r] == target {
				r++
			}
			return []int{l + 1, r - 1}
		} else if nums[mid] < target {
			for mid+1 < right && nums[mid] == nums[mid+1] {
				mid++
			}
			left = mid + 1
		} else {
			for mid-1 > left && nums[mid] == nums[mid-1] {
				mid--
			}
			right = mid - 1
		}
	}
	return ans
}

func Test_searchRange(t *testing.T) {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchRange([]int{5, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 8, 10, 10, 10, 10, 10, 10, 10, 10, 10, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 8}, 8))
	fmt.Println(searchRange([]int{}, 6))
}

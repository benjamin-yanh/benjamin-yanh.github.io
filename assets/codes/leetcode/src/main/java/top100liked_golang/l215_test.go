package top100liked_golang

import "sort"

func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k-1]
}

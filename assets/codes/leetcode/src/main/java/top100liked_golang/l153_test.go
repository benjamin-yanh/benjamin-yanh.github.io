package top100liked_golang

import (
	"fmt"
	"testing"
)

func findMin(nums []int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	if n == 1 {
		return nums[0]
	}
	left, right := 0, n-1
	var mid = -1
	for left <= right {
		mid = left + (right-left)/2
		if nums[left] <= nums[mid] && nums[left] <= nums[right] {
			minValue := nums[left]
			for i := left; i >= 0; i-- {
				if nums[i] > minValue {
					break
				}
				minValue = nums[i]
			}
			return minValue
		} else if nums[left] <= nums[mid] && nums[right] < nums[mid] {
			left = mid + 1
		} else if nums[mid] < nums[left] && nums[mid] < nums[right] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return nums[mid]
}

func Test_findMin(t *testing.T) {
	fmt.Println(findMin([]int{1, 2, 3, 4, 5}))       // 1
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2})) // 0
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))       // 1
	fmt.Println(findMin([]int{11, 13, 15, 17}))      // 11
	fmt.Println(findMin([]int{3, 1, 2}))             // 1
	fmt.Println(findMin([]int{2, 1}))                // 1
	fmt.Println(findMin([]int{2, 3, 0}))             // 0
	fmt.Println(findMin([]int{1, 2, 3, 4, 0}))       // 0
}

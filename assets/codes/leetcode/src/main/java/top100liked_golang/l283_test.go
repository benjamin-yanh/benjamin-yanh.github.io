package top100liked_golang

import "testing"

func moveZeroes2(nums []int) {

	zeroNums := make([]int, len(nums))
	index := 0
	for _, num := range nums {
		if num == 0 {
			continue
		}
		zeroNums[index] = num
		index++
	}
	copy(nums, zeroNums)
}

func moveZeroes(nums []int) {
	left, right, n := 0, 0, len(nums)

	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
	return
}

func Test_moveZeroes(t *testing.T) {
	moveZeroes([]int{2, 1, 0, 3, 12})
	//moveZeroes([]int{0, 1, 0, 3, 12})
}

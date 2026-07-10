package top100liked_golang

import (
	"fmt"
	"testing"
)

//func productExceptSelf(nums []int) []int {
//	zero := 0
//	for _, num := range nums {
//		if num == 0 {
//			zero++
//		}
//	}
//
//	if zero > 1 {
//		return make([]int, len(nums))
//	}
//
//	if zero == 1 {
//		val := 1
//		index := 0
//		for i, num := range nums {
//			if num == 0 {
//				index = i
//			} else {
//				val *= num
//				nums[i] = 0
//			}
//		}
//		nums[index] = val
//		return nums
//	}
//
//	val := 1
//	for i := 1; i < len(nums); i++ {
//		val *= nums[i]
//	}
//
//	temp := nums[0]
//	nums[0] = val
//	for i := 1; i < len(nums); i++ {
//		nums[i] = val / nums[i] * temp
//	}
//	return nums
//}

//func productExceptSelf(nums []int) []int {
//	left := make([]int, len(nums))
//	right := make([]int, len(nums))
//	left[0] = 1
//	for i := 1; i < len(nums); i++ {
//		left[i] = nums[i-1] * left[i-1]
//	}
//
//	right[len(nums)-1] = 1
//	for i := len(nums) - 2; i >= 0; i-- {
//		right[i] = nums[i+1] * right[i+1]
//	}
//
//	for i := 0; i < len(nums); i++ {
//		nums[i] = left[i] * right[i]
//	}
//
//	return nums
//}

func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	length := len(nums)
	ans[0] = 1
	for i := 1; i < length; i++ {
		ans[i] = nums[i-1] * ans[i-1]
	}

	right := 1
	for i := length - 1; i >= 0; i-- {
		ans[i] = ans[i] * right
		right *= nums[i]
	}

	return ans
}

func Test_productExceptSelf(t *testing.T) {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
}

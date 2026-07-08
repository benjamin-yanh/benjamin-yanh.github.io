package top100liked_golang

import (
	"fmt"
	"testing"
)

func trap(height []int) int {
	left, right, leftMax, rightMax := 0, len(height)-1, 0, 0
	sum := 0

	for left < right {
		leftMax = max(height[left], leftMax)
		rightMax = max(height[right], rightMax)

		if height[left] < height[right] {
			sum += leftMax - height[left]
			left++
		} else {
			sum += rightMax - height[right]
			right--
		}
	}

	return sum
}

func Test_trap(t *testing.T) {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

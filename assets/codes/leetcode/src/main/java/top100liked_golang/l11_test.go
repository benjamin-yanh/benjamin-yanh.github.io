package top100liked_golang

import (
	"fmt"
	"testing"
)

func maxArea(height []int) int {
	start := 0
	end := len(height) - 1
	maxValue := 0

	for start < end {
		minValue, _ := min(height[start], height[end])
		maxValue = max(maxValue, minValue*(end-start))
		if height[start] < height[end] {
			start++
		} else {
			end--
		}
	}

	return maxValue
}

func Test_maxArea(t *testing.T) {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 1}))
	fmt.Println(maxArea([]int{1, 2}))
}

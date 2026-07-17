package top100liked_golang

import (
	"fmt"
	"testing"
)

func largestRectangleArea(heights []int) int {
	n := len(heights)
	left := make([]int, n)
	right := make([]int, n)
	stack := []int{}
	for i := 0; i < n; i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	stack = []int{}
	for j := n - 1; j >= 0; j-- {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[j] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			right[j] = n
		} else {
			right[j] = stack[len(stack)-1]
		}
		stack = append(stack, j)
	}
	maxV := 0
	for i := 0; i < n; i++ {
		maxV = max(maxV, (right[i]-left[i]-1)*heights[i])
	}
	return maxV
}

func Test_largestRectangleArea(t *testing.T) {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}

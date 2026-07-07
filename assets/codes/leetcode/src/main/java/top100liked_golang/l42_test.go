package top100liked_golang

import (
	"fmt"
	"testing"
)

func trap(height []int) int {
	left, right, n := 0, 1, len(height)
	total := 0
	sum := 0
	{
		for right < n {
			if height[right] > height[left] {
				total += sum
				sum = 0
				left = right
				right++
				continue
			}
			sum += height[left] - height[right]
			right++
		}
	}

	return total
}

func Test_trap(t *testing.T) {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

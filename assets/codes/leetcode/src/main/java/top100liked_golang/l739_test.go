package top100liked_golang

import (
	"fmt"
	"testing"
)

func dailyTemperatures(temperatures []int) []int {
	stack := []int{}
	result := make([]int, len(temperatures))

	for i, ele := range temperatures {
		if len(stack) == 0 {
			stack = append(stack, i)
			continue
		}
		for len(stack) > 0 && ele > temperatures[stack[len(stack)-1]] {
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[last] = i - last
		}
		stack = append(stack, i)
	}
	return result
}

func Test_dailyTemperatures(t *testing.T) {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
	fmt.Println(dailyTemperatures([]int{30, 40, 50, 60}))
	fmt.Println(dailyTemperatures([]int{30, 40, 50, 60}))
	fmt.Println(dailyTemperatures([]int{30, 60, 90}))
}

package top100liked_golang

import (
	"fmt"
	"testing"
)

func maxSlidingWindow(nums []int, k int) []int {
	result := []int{}
	queue := []int{}

	for i, num := range nums {
		// 右边进
		for len(queue) > 0 && nums[queue[len(queue)-1]] < num {
			// 如果队列中的元素都已经小于当前元素的就不用考虑了，num 将会进入窗口了
			queue = queue[:len(queue)-1]
		}
		// 如果没有元素那么，当前就是最大的元素
		queue = append(queue, i) // num 进入单调窗口
		// 左边出，最小的元素超出了窗口就移除
		left := i - k + 1
		if queue[0] < left {
			queue = queue[1:]
		}
		// 只有最左元素大于等于窗口大小了 才有效
		if left >= 0 {
			result = append(result, nums[queue[0]])
		}
	}

	return result
}

func Test_maxSlidingWindow(t *testing.T) {
	fmt.Println(maxSlidingWindow([]int{1}, 1))
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}

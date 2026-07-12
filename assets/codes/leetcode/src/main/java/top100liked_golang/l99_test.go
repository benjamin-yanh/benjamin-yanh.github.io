package top100liked_golang

import "testing"

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ans := []int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		ans = append(ans, queue[len(queue)-1].Val)
		length := len(queue)
		for i := 0; i < length; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[length:]
	}
	return ans
}

func Test_rightSideView(t *testing.T) {

}

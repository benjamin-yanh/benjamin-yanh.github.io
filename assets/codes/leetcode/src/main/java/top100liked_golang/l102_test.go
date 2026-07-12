package top100liked_golang

import (
	"fmt"
	"testing"
)

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ans := [][]int{}
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) > 0 {
		rows := []int{}
		length := len(queue)
		for i := 0; i < length; i++ {
			rows = append(rows, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		ans = append(ans, rows)
		queue = queue[length:]
	}
	return ans
}

func printlnTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	printlnTree(root.Left)
	printlnTree(root.Right)
}

func Test_levelOrder(t *testing.T) {
	root := &TreeNode{Val: 3,
		Left: &TreeNode{Val: 9,
			Left:  &TreeNode{Val: 5},
			Right: &TreeNode{Val: 7},
		},
		Right: &TreeNode{Val: 20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 17},
		},
	}
	printlnTree(root)
	//levelOrder(root)
}

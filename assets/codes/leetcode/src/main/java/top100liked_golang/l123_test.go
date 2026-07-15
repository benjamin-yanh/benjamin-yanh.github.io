package top100liked_golang

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxValue = -1001

func maxPathSum(root *TreeNode) int {
	maxValue = root.Val
	maxPathSum2(root)
	return maxValue
}

func maxPathSum2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxPathSumNode(root, root)
	maxPathSum2(root.Left)
	maxPathSum2(root.Right)
	return maxValue
}

func maxPathSumNode(node, root *TreeNode) int {
	if node == nil {
		return 0
	}
	left := maxPathSumNode(node.Left, root)
	right := maxPathSumNode(node.Right, root)
	var val int
	if root == node {
		val = max(node.Val+left, node.Val+right, node.Val, node.Val+left+right)
	} else {
	}
	if val > maxValue {
		maxValue = val
	}
	return val
}

func Test_maxPathSum(t *testing.T) {
	//root := &TreeNode{Val: 15,
	//	Left: &TreeNode{Val: 9,
	//		Left: &TreeNode{Val: 8,
	//			Left:  &TreeNode{Val: 6},
	//			Right: &TreeNode{Val: 7},
	//		},
	//		Right: &TreeNode{Val: 10,
	//			Right: &TreeNode{Val: 11},
	//		},
	//	},
	//	Right: &TreeNode{Val: 16,
	//		Right: &TreeNode{Val: 17}}}

	root := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3}}

	//root := &TreeNode{Val: 0}
	val := maxPathSum(root)
	fmt.Printf("maxPathSum root = %v\n", val)
}

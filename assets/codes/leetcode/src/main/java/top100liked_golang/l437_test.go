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

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	res := rootSum(root, targetSum)
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)

	return res
}

func rootSum(root *TreeNode, targetSum int) (res int) {
	if root == nil {
		return
	}
	val := root.Val
	if val == targetSum {
		res++
	}
	res += rootSum(root.Left, targetSum-val)
	res += rootSum(root.Right, targetSum-val)
	return
}

func preOrderSum(root *TreeNode, currentSum int) int {
	if currentSum == 0 {
		return 1
	}
	if root == nil {
		return 0
	}
	left := preOrderSum(root.Left, currentSum-root.Val)
	right := preOrderSum(root.Right, currentSum-root.Val)
	return left + right
}

//func pathLeftSum(root *TreeNode, targetSum int) int {
//	if targetSum == 0 {
//		return 1
//	}
//	if root == nil {
//		return 0
//	}
//	return pathLeftSum(root.Left, targetSum-root.Val)
//
//}
//
//func pathRightSum(root *TreeNode, targetSum int) int {
//
//	if targetSum == 0 {
//		return 1
//	}
//	if root == nil {
//		return 0
//	}
//
//	return pathRightSum(root.Right, targetSum-root.Val)
//}

func Test_pathSum(t *testing.T) {
	root := &TreeNode{Val: 10,
		Left: &TreeNode{Val: 5,
			Left: &TreeNode{Val: 3,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: -2},
			},
			Right: &TreeNode{Val: 2,
				Right: &TreeNode{Val: 1},
			},
		},
		Right: &TreeNode{Val: -3,
			Right: &TreeNode{Val: 11}}}

	fmt.Println(pathSum(root, 8))
}

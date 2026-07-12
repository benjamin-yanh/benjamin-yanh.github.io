package top100liked_golang

import (
	"fmt"
	"testing"
)

//func flatten(root *TreeNode) {
//	if root == nil {
//		return
//	}
//	ans := []int{}
//	preorder(root, &ans)
//	root.Left = nil
//	root.Right = nil
//	right := root
//	for i := 1; i < len(ans); i++ {
//		right.Right = &TreeNode{Val: ans[i], Left: nil, Right: nil}
//		right = right.Right
//	}
//	return
//}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	var x *TreeNode = nil
	if root.Left != nil {
		x = root.Left
	}
	var t *TreeNode = nil
	if x != nil {
		t = x.Right
	}
	if x != nil {
		x.Right = root
	}
	root.Left = t
	flatten(root.Left)
	flatten(root.Right)
	return
}

func preorder(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	*ans = append(*ans, root.Val)
	preorder(root.Left, ans)
	preorder(root.Right, ans)
}

func printlnFlatten(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	printlnFlatten(root.Left)
	printlnFlatten(root.Right)
}

func Test_flatten(t *testing.T) {
	root := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 6}}}
	flatten(root)
	printlnFlatten(root)
}

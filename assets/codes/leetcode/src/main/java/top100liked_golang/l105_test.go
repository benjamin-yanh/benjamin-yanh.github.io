package top100liked_golang

import (
	"fmt"
	"testing"
)

var inorderMap = make(map[int]int)
var preorderMap = make(map[int]int)

//func buildTree(preorder []int, inorder []int) *TreeNode {
//	root := &TreeNode{Val: preorder[0]}
//	for i, ele := range inorder {
//		inorderMap[ele] = i
//	}
//	for i, ele := range preorder {
//		preorderMap[ele] = i
//	}
//
//	buildTreeNode(preorder[1:], inorder, inorderMap[root.Val], 0, len(inorder), root)
//	return root
//}

func buildTreeNode(preorder, inorder []int, inorderIndex, inorderStart, inorderEnd int, root *TreeNode) {
	if len(preorder) <= 0 {
		return
	}

	{
		// left
		val := preorder[0]
		rootIndex := inorderMap[val]
		if rootIndex < inorderIndex && rootIndex >= inorderStart {
			// 存在左节点，且为根节点
			root.Left = &TreeNode{Val: val}
			buildTreeNode(preorder[1:], inorder, inorderMap[val], inorderStart, inorderIndex, root.Left)
		}
	}
	{
		// right
		for index, val := range preorder {
			rootIndex := inorderMap[val]
			if rootIndex < inorderEnd && rootIndex >= inorderIndex {
				// 存在右节点，且为根节点
				root.Right = &TreeNode{Val: val}
				buildTreeNode(preorder[index+1:], inorder, inorderMap[val], inorderIndex, inorderEnd, root.Right)
				break
			}
		}
	}
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) <= 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])

	return root
}

func Test_buildTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	fmt.Println(buildTree(preorder, inorder))
}

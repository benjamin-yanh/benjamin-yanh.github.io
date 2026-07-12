package top100liked_golang

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	val := dsfDiameterOfBinaryTree(root.Left) + dsfDiameterOfBinaryTree(root.Right)
	left := diameterOfBinaryTree(root.Left)
	right := diameterOfBinaryTree(root.Right)
	return max(left, right, val)
}

func dsfDiameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := 1 + dsfDiameterOfBinaryTree(root.Left)
	right := 1 + dsfDiameterOfBinaryTree(root.Right)
	return max(left, right)
}

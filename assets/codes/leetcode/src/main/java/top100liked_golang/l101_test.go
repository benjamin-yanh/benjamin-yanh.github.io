package top100liked_golang

func isSymmetric(root *TreeNode) bool {

	return isSymmetricLCheck(root.Left, root.Right)
}

func isSymmetricLCheck(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val && isSymmetricLCheck(left.Left, right.Right) && isSymmetricLCheck(left.Right, right.Left)
}

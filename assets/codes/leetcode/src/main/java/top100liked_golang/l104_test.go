package top100liked_golang

func maxDepth(root *TreeNode) int {
	return inorderMaxDepth(root, 0)
}
func inorderMaxDepth(root *TreeNode, ans int) int {
	if root == nil {
		return ans
	}

	r := inorderMaxDepth(root.Left, ans+1)
	l := inorderMaxDepth(root.Right, ans+1)
	return max(r, l)
}

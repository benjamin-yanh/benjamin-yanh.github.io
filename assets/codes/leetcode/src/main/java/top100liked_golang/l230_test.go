package top100liked_golang

func kthSmallest(root *TreeNode, k int) int {
	ans := []int{}
	inorderKthSmallest(root, &ans)
	return ans[k-1]
}

func inorderKthSmallest(root *TreeNode, ans *[]int) {
	if root == nil {

		return
	}
	inorderKthSmallest(root.Left, ans)
	*ans = append(*ans, root.Val)
	inorderKthSmallest(root.Right, ans)
}

package top100liked_golang

import "testing"

func isValidBST(root *TreeNode) bool {
	ans := []int{}
	preOrderBST(root, &ans)
	for i := 0; i < len(ans)-1; i++ {
		if ans[i] > ans[i+1] {
			return false
		}
	}
	return true
}

func preOrderBST(root *TreeNode, asn *[]int) {
	if root == nil {
		return
	}
	preOrderBST(root.Left, asn)
	preOrderBST(root.Right, asn)
	*asn = append(*asn, root.Val)
}

func Test_isValidBST(t *testing.T) {

}

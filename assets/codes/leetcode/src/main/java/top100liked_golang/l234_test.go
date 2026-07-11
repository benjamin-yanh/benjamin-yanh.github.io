package top100liked_golang

func isPalindrome(head *ListNode) bool {
	ans := []int{}
	for head != nil {
		ans = append(ans, head.Val)
		head = head.Next
	}
	for i := 0; i < len(ans); i++ {
		if ans[i] != ans[len(ans)-i-1] {
			return false
		}
	}

	return true
}

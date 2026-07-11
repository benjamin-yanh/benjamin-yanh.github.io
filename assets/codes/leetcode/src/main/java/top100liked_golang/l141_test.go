package top100liked_golang

import "testing"

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	temp := head.Next
	for temp != nil && temp.Next != nil && head != nil {
		if temp == head {
			return true
		}
		temp = temp.Next.Next
		head = head.Next
	}

	return false
}

func Test_hasCycle(t *testing.T) {

}

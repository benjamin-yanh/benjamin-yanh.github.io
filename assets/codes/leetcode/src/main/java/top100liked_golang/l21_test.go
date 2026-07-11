package top100liked_golang

import "testing"

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{Val: -1, Next: nil}
	temp := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			temp.Next = &ListNode{Val: list1.Val, Next: nil}
			list1 = list1.Next
		} else {
			temp.Next = &ListNode{Val: list2.Val, Next: nil}
			list2 = list2.Next
		}
		temp = temp.Next
	}
	if list1 != nil {
		temp.Next = list1
	}
	if list2 != nil {
		temp.Next = list2
	}
	return dummy.Next
}

func Test_mergeTwoLists(t *testing.T) {

}

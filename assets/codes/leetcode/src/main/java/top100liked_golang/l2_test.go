package top100liked_golang

import "testing"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	temp := dummy
	val := 0
	for l1 != nil || l2 != nil {
		v1 := 0
		v2 := 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		sum := v1 + v2 + val
		temp.Next = &ListNode{Val: sum % 10}
		val = sum / 10
		temp = temp.Next
	}
	if val != 0 {
		temp.Next = &ListNode{Val: val}
	}
	return dummy.Next
}

func Test_addTwoNumbers(t *testing.T) {

}

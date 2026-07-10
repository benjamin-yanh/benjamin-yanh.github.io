package top100liked_golang

import "testing"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func Test_reverseList(t *testing.T) {
	reverseList(&ListNode{Val: 1, Next: &ListNode{2, &ListNode{Val: 3, Next: &ListNode{5, nil}}}})
}

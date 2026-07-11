package top100liked_golang

import "testing"

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	maps := make(map[*ListNode]int)
	for head != nil {
		if _, ok := maps[head]; ok {
			return head
		}
		maps[head] = 1
		head = head.Next
	}

	return nil

}

func Test_detectCycle(t *testing.T) {

}

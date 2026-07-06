package top100liked_golang

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}

	for l1 != nil && l2 != nil {
		next := dummy.Next
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
		node := &ListNode{}
		node.Next = next
		node.Val = v1 + v2
		dummy.Next = node
	}

	v := 0
	node := dummy.Next
	for node != nil {
		v1 := (node.Val + v) / 10
		node.Val = (node.Val + v) % 10
		v = v1
		node = node.Next
	}

	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}

package top100liked_golang

import (
	"sort"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	list := []*ListNode{}
	for head != nil {
		list = append(list, head)
		head = head.Next
	}

	sort.Slice(list, func(i int, j int) bool {
		return list[i].Val < list[j].Val
	})
	for i := 0; i < len(list)-1; i++ {
		list[i].Next = list[i+1]
	}
	list[len(list)-1].Next = nil
	return list[0]
}

func Test_sortList(t *testing.T) {
	node := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
				}}}}
	list := sortList(node)
	for list != nil {
		t.Log(list.Val)
		list = list.Next
	}
}

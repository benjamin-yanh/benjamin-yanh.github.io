package top100liked_golang

import (
	"fmt"
	"testing"
)

func swapPairs(head *ListNode) *ListNode {
	ans := []*ListNode{}
	temp := head
	for temp != nil {
		ans = append(ans, temp)
		if len(ans) == 2 {
			ans[0].Val, ans[1].Val = ans[1].Val, ans[0].Val
			ans = ans[:0]
		}
		temp = temp.Next
	}
	return head
}

func Test_swapPairs(t *testing.T) {
	fmt.Println(swapPairs(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}}))
}

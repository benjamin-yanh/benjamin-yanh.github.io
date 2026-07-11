package top100liked_golang

import (
	"fmt"
	"testing"
)

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: -1, Next: nil}
	dummy.Next = head
	temp := dummy
	stack := []*ListNode{}
	length := 0
	for temp != nil {
		stack = append(stack, temp)
		temp = temp.Next
		length++
	}

	stack[length-n-1].Next = stack[length-n-1].Next.Next

	return dummy.Next
}

func Test_removeNthFromEnd(t *testing.T) {
	fmt.Println(removeNthFromEnd(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}, 2))
	fmt.Println(removeNthFromEnd(&ListNode{Val: 1, Next: nil}, 1))

}

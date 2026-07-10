package top100liked_golang

import "testing"

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	maps := map[*ListNode]int{}
	for headA != nil {
		maps[headA] = 0
		headA = headA.Next
	}
	for headB != nil {
		if _, ok := maps[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

func Test_getIntersectionNode(t *testing.T) {

}

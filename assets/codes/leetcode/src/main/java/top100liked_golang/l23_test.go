package top100liked_golang

import "sort"

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	ints := []int{}
	for _, list := range lists {
		for list != nil {
			ints = append(ints, list.Val)
			list = list.Next
		}
	}
	sort.Ints(ints)
	dummy := &ListNode{}
	cur := dummy
	for _, int := range ints {
		cur.Next = &ListNode{Val: int}
		cur = cur.Next
	}
	return dummy.Next

}

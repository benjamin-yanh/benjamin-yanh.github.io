package top100liked_golang

import (
	"fmt"
	"testing"
)

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	oldRandomMaps := make(map[int]*Node)

	oldNodesIndex := make(map[*Node]int)

	curNodes := []*Node{}

	dummy := &Node{Val: -1}
	cur := dummy
	index := 0
	for head != nil {
		cur.Next = &Node{Val: head.Val}
		oldRandomMaps[index] = head.Random
		oldNodesIndex[head] = index
		curNodes = append(curNodes, cur.Next)
		cur = cur.Next
		head = head.Next
		index++
	}

	for key, random := range oldRandomMaps {
		if random == nil {
			continue
		}
		randomIndex := oldNodesIndex[random]
		randomNode := curNodes[randomIndex]
		curNodes[key].Random = randomNode
	}

	return dummy.Next

}

func Test_copyRandomList(t *testing.T) {
	node0 := &Node{Val: 7}
	node1 := &Node{Val: 13}
	node2 := &Node{Val: 11}
	node3 := &Node{Val: 10}
	node4 := &Node{Val: 1}

	node0.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	node1.Random = node0
	node2.Random = node4
	node3.Random = node2
	node4.Random = node0

	dummy := copyRandomList(node0)
	for dummy != nil {
		if dummy.Random != nil {
			fmt.Println(dummy.Val, ",", dummy.Random.Val)
		} else {
			fmt.Println(dummy.Val)
		}
		dummy = dummy.Next
	}
}

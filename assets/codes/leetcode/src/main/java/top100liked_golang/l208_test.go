package top100liked_golang

import (
	"slices"
)

type Trie struct {
	Value []string
	Next  []*Trie
}

func Constructor() Trie {
	return Trie{
		Value: []string{},
		Next:  make([]*Trie, 128),
	}
}

func (this *Trie) Insert(word string) {
	cur := this
	for _, ele := range word {
		if cur.Next[ele] == nil {
			cur.Next[ele] = &Trie{
				Value: []string{},
				Next:  make([]*Trie, 128),
			}
		}
		cur = cur.Next[ele]
	}
	cur.Value = append(cur.Value, word)
}

func (this *Trie) Search(word string) bool {
	cur := this
	for _, ele := range word {
		if cur.Next[ele] == nil {
			return false
		}
		cur = cur.Next[ele]
	}
	if cur == nil {
		return false
	}
	return slices.Contains(cur.Value, word)
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for _, ele := range prefix {
		if cur.Next[ele] == nil {
			return false
		}
		cur = cur.Next[ele]
	}
	return true
}

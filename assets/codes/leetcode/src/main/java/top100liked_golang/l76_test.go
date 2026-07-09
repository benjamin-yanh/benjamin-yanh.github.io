package top100liked_golang

import (
	"fmt"
	"testing"
)

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	mark_t := [128]int{}
	for _, ele := range t {
		mark_t[ele]++
	}

	mark_s := [128]int{}
	oven_s := [128]int{}

	left, right, minLeft, minRight := 0, 0, -1, len(s)

	for right < len(s) {

		for ; mark_s != mark_t && right < len(s); right++ {
			v := s[right]
			if mark_t[v] == 0 {
				continue
			}
			if mark_s[v] < mark_t[v] {
				mark_s[v]++
			} else {
				oven_s[v]++
			}
		}
		if mark_t == mark_s {
			minLeft, minRight = minScope(left, right, minLeft, minRight)
		}

		for ; mark_s == mark_t && left < right; left++ {
			v := s[left]
			if oven_s[v] > 0 {
				oven_s[v]--
			} else if mark_t[v] > 0 {
				mark_s[v]--
			}
			if mark_t == mark_s {
				minLeft, minRight = minScope(left+1, right, minLeft, minRight)
			}
		}
	}

	if minLeft < 0 {
		return ""
	}

	return s[minLeft:minRight]
}

func minScope(left, right, minLeft, minRight int) (int, int) {
	if right-left < minRight-minLeft {
		return left, right
	}
	return minLeft, minRight
}

func Test_minWindow(t *testing.T) {
	fmt.Println(minWindow("ab", "b"))                  //b
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))     //BANC
	fmt.Println(minWindow("a", "a"))                   // a
	fmt.Println(minWindow("cabwefgewcwaefgcf", "cae")) // cwae
}

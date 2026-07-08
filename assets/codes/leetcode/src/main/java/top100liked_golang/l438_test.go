package top100liked_golang

import (
	"fmt"
	"testing"
)

func findAnagrams(s string, p string) []int {
	var result []int
	var parr [26]int
	for _, e := range p {
		parr[e-'a']++
	}

	left := 0
	right := 0
	var sarr [26]int
	for ; right < len(s); right++ {
		if right < len(p) {
			sarr[s[right]-'a']++
		} else {
			sarr[s[right]-'a']++
			sarr[s[left]-'a']--
			left++
		}
		if sarr == parr {
			result = append(result, left)
		}
	}

	return result
}

func Test_findAnagrams(t *testing.T) {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
	fmt.Println(findAnagrams("a", "a"))
}

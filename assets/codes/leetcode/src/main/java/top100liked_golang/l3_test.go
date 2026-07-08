package top100liked_golang

import (
	"fmt"
	"testing"
)

func lengthOfLongestSubstring(s string) int {
	maps := make(map[byte]int)
	size := 0
	left := 0
	right := 0

	for right < len(s) {
		for left < right {
			if _, ok := maps[s[right]]; ok {
				maps[s[left]]--
				if maps[s[left]] == 0 {
					delete(maps, s[left])
				}
			} else {
				break
			}
			left++
		}
		maps[s[right]]++
		right++
		size = max(size, len(maps))
	}
	return size
}

func Test_lengthOfLongestSubstring(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring(""))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

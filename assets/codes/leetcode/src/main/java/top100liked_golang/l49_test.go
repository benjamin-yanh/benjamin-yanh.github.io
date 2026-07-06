package top100liked_golang

import (
	"sort"
	"strings"
	"testing"
)

func groupAnagrams(strs []string) [][]string {
	maps := make(map[string][]string)
	for _, str := range strs {
		strs := strings.Split(str, "")
		sort.Strings(strs)
		val := strings.Join(strs, "")
		if _, ok := maps[val]; !ok {
			maps[val] = []string{}
		}
		maps[val] = append(maps[val], str)
	}
	var result [][]string
	for _, ele := range maps {
		result = append(result, ele)
	}
	return result
}

func groupAnagrams2(strs []string) [][]string {
	maps := make(map[[26]int][]string)
	for _, str := range strs {
		cnt := [26]int{}
		for _, ele := range str {
			cnt[ele-'a']++
		}
		maps[cnt] = append(maps[cnt], str)
	}
	var result [][]string
	for _, ele := range maps {
		result = append(result, ele)
	}
	return result
}

func Test_groupAnagrams(t *testing.T) {

}

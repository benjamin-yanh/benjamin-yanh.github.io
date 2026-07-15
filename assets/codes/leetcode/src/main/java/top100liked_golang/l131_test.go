package top100liked_golang

import (
	"fmt"
	"slices"
	"testing"
)

func partition(s string) [][]string {
	result := [][]string{}
	n := len(s)
	findPartition(s, 0, 0, n, []string{}, &result)
	return result
}

func findPartition(s string, i, start, n int, path []string, result *[][]string) {
	if i == n {
		*result = append(*result, slices.Clone(path))
		return
	}

	if i < n-1 {
		findPartition(s, i+1, start, n, path, result)
	}
	substr := s[start : i+1]
	if isPartition(substr) { // 如果是回文就判断子串分割方案
		path = append(path, substr)
		findPartition(s, i+1, i+1, n, path, result)
		path = path[:len(path)-1]
	}
}

//if isPartition(val) {
//	*result = append(*result, val)
//}

func isPartition(s string) bool {

	length := len(s)

	for i := 0; i < length/2; i++ {
		if s[i] != s[length-i-1] {
			return false
		}
	}
	return true
}

func Test_partition(t *testing.T) {
	val := "aab"
	fmt.Println(partition(val))
}

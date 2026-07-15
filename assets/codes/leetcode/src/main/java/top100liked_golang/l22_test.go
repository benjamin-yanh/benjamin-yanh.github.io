package top100liked_golang

import (
	"fmt"
	"testing"
)

func generateParenthesis(n int) []string {
	result := []string{}

	findGenerateParenthesis(0, 0, n, "", &result)

	return result
}

func findGenerateParenthesis(left, right, n int, ans string, result *[]string) {
	if len(ans) == 2*n {
		*result = append(*result, ans)
		return
	}
	if left < n {
		findGenerateParenthesis(left+1, right, n, ans+"(", result)
	}
	if left > right {
		findGenerateParenthesis(left, right+1, n, ans+")", result)
	}
}

func Test_generateParenthesis(t *testing.T) {
	result := generateParenthesis(8)
	for _, item := range result {
		fmt.Println(item)
	}
}

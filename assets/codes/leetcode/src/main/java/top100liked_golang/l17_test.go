package top100liked_golang

import "testing"

var keys = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	results := []string{}
	findLetterCombinations(digits, "", &results)
	return results
}
func findLetterCombinations(digits, ans string, result *[]string) {
	if len(digits) == 0 {
		*result = append(*result, ans)
		return
	}
	number := digits[0] - '0'
	abc := keys[number]
	for i := 0; i < len(abc); i++ {
		findLetterCombinations(digits[1:], ans+string(abc[i]), result)
	}
}
func Test_letterCombinations(t *testing.T) {

}

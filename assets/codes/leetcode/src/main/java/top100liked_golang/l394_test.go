package top100liked_golang

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func decodeString(s string) string {
	stack := make([]string, 0)
	for _, ele := range s {
		if ele == ']' {
			sb := strings.Builder{}
			substr := ""

			for stack[len(stack)-1] != "[" {
				substr = stack[len(stack)-1] + substr
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
			numstr := ""
			for len(stack)-1 >= 0 && stack[len(stack)-1] >= "0" && stack[len(stack)-1] <= "9" {
				numstr = stack[len(stack)-1] + numstr
				stack = stack[:len(stack)-1]
			}
			number, _ := strconv.Atoi(numstr)
			for j := 0; j < number; j++ {
				sb.WriteString(substr)
			}
			stack = append(stack, sb.String())
		} else {
			stack = append(stack, string(ele))
		}
	}

	return strings.Join(stack, "")
}

func Test_decodeString(t *testing.T) {
	fmt.Println(decodeString("a") == "a")
	fmt.Println(decodeString("3[a]") == "aaa")
	fmt.Println(decodeString("3[a2[c]]") == "accaccacc")
	fmt.Println(decodeString("2[abc]3[cd]ef") == "abcabccdcdcdef")
	fmt.Println(decodeString("abc3[cd]xyz") == "abccdcdcdxyz")
}

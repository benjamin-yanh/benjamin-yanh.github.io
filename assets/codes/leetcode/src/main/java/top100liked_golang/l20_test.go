package top100liked_golang

func isValid(s string) bool {
	ans := []string{}
	for i := 0; i < len(s); i++ {
		if s[i] == '{' || s[i] == '(' || s[i] == '[' {
			ans = append(ans, string(s[i]))
			continue
		}
		if len(ans) == 0 {
			return false
		}
		v := ans[len(ans)-1]
		ans = ans[:len(ans)-1]
		if s[i] == '}' && v != "{" {
			return false
		}
		if s[i] == ']' && v != "[" {
			return false
		}

		if s[i] == ')' && v != "(" {
			return false
		}


	}
	return len(ans) == 0
}

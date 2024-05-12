package main

import "fmt"

func isMatch(s string, p string) bool {
	if p == ".*" {
		return true
	}
	if s == p {
		return true
	}

	if len(s) == 0 || len(p) == 0 {
		return false
	}
	if p[0] == '.' {
		return isMatch(s[1:], p[1:])
	}

	if len(p) == 1 {
		return false
	}

	for i := 0; i < len(s); i++ {
		if p[0] == s[i] || p[0] == '.' {
			if p[1] == '*' {
				return isMatch(s[i+1:], p[2:]) || isMatch(s[i+1:], p) || isMatch(s[i:], p[2:])
			} else {
				return isMatch(s[i+1:], p[1:])
			}
		}
	}

	return false
}

func main() {
	s, p := "aa", "a"
	fmt.Println(isMatch(s, p))
	s, p = "aa", "a*"
	fmt.Println(isMatch(s, p))
	s, p = "ab", ".*"
	fmt.Println(isMatch(s, p))

}

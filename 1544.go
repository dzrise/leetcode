package main

import (
	"fmt"
	"strings"
)

func makeGood(s string) string {
	if len(s) < 2 {
		return s
	}

	i := 0
	for {
		fmt.Println(i)
		if i >= len(s)-1 {
			break
		}
		if s[i] != s[i+1] &&
			(string(s[i]) == strings.ToUpper(string(s[i+1])) || strings.ToUpper(string(s[i])) == string(s[i+1])) {
			news := s[:i] + s[i+2:]
			s = news
			if i > 0 {
				i--
			}

		} else {
			i++
		}
	}

	return s
}

func main() {
	fmt.Println(makeGood("leEeetcode"))
	//fmt.Println(makeGood("abBAcC"))
}

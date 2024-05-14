//22. Generate Parentheses
//Medium
//Topics
//Companies
//Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
//
//
//
//Example 1:
//
//Input: n = 3
//Output: ["((()))","(()())","(())()","()(())","()()()"]
//Example 2:
//
//Input: n = 1
//Output: ["()"]
//
//
//Constraints:

package main

import "fmt"

func _generateParenthesis(left, right int, s *string, sl *[]string) {

	if left == 0 && right == 0 {
		*sl = append(*sl, *s)
	}

	if left > 0 {
		*s += "("
		_generateParenthesis(left-1, right, &*s, &*sl)
	}

	if right > 0 && left < right {
		*s += ")"
		_generateParenthesis(left, right-1, &*s, &*sl)
	}

	if len(*s) > 0 {
		*s = (*s)[:len(*s)-1]
	}

}
func generateParenthesis(n int) []string {
	sl := make([]string, 0)
	var s string

	_generateParenthesis(n, n, &s, &sl)

	return sl

}

func main() {
	fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis(1))
}

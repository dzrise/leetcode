// 20. Valid Parentheses
// Easy
// Topics
// Companies
// Hint
// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
// An input string is valid if:
//
// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.
//
// Example 1:
//
// Input: s = "()"
// Output: true
// Example 2:
//
// Input: s = "()[]{}"
// Output: true
// Example 3:
//
// Input: s = "(]"
// Output: false
//
// Constraints:
//
// 1 <= s.length <= 104
// s consists of parentheses only '()[]{}'.
package main

import (
	"fmt"
)

func isValid(s string) bool {
	need := make([]byte, 0, len(s)/2)

	for i := 0; i < len(s); i++ {
		char := s[i]
		switch char {
		case '(':
			need = append(need, ')')
		case '[':
			need = append(need, ']')
		case '{':
			need = append(need, '}')
		default:
			last := len(need) - 1
			if last < 0 || need[last] != char {
				return false
			}
			need = need[:last]

		}
	}

	return len(need) == 0

}

//func isValid(s string) bool {
//	i := len(s)
//	for i > 0 {
//		s = strings.Replace(s, "()", "", -1)
//		s = strings.Replace(s, "[]", "", -1)
//		s = strings.Replace(s, "{}", "", -1)
//
//		if len(s) == 0 {
//			return true
//		}
//
//		i--
//	}
//	return false
//}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([])"))
}

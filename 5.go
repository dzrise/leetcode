//Given a string s, return the longest
//palindromic
//
//substring
//in s.
//
//
//
//Example 1:
//
//Input: s = "babad"
//Output: "bab"
//Explanation: "aba" is also a valid answer.
//Example 2:
//
//Input: s = "cbbd"
//Output: "bb"
//
//
//Constraints:
//
//1 <= s.length <= 1000
//s consist of only digits and English letters.

package main

import (
	"fmt"
)

func getPalindrom(s string, left int, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	if len(s) == 2 {
		if s[0] == s[1] {
			return s
		}
	}
	res := s[0:1]
	for i := 0; i < len(s)-1; i++ {
		s1, s2 := getPalindrom(s, i, i), getPalindrom(s, i, i+1)
		if len(s1) > len(res) {
			res = s1
		}
		if len(s2) > len(res) {
			res = s2
		}
	}

	return res

}

func main() {
	s := "babad"
	fmt.Println(longestPalindrome(s))
	s = "cbbd"
	fmt.Println(longestPalindrome(s))
	s = "ccd"
	fmt.Println(longestPalindrome(s))
}

// 10. Regular Expression Matching
// Attempted
// Hard
// Topics
// Companies
// Given an input string s and a pattern p, implement regular expression matching with support for '.' and '*' where:
//
// '.' Matches any single character.​​​​
// '*' Matches zero or more of the preceding element.
// The matching should cover the entire input string (not partial).
//
// Example 1:
//
// Input: s = "aa", p = "a"
// Output: false
// Explanation: "a" does not match the entire string "aa".
// Example 2:
//
// Input: s = "aa", p = "a*"
// Output: true
// Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".
// Example 3:
//
// Input: s = "ab", p = ".*"
// Output: true
// Explanation: ".*" means "zero or more (*) of any character (.)".
//
// Constraints:
//
// 1 <= s.length <= 20
// 1 <= p.length <= 20
// s contains only lowercase English letters.
// p contains only lowercase English letters, '.', and '*'.
// It is guaranteed for each appearance of the character '*', there will be a previous valid character to match.
package main

import "fmt"

func isMatch(s string, p string) bool {
	rows := len(s)
	cols := len(p)

	if rows == 0 && cols == 0 {
		return true
	}

	if cols == 0 {
		return false
	}

	dp := make([][]int, rows+1)

	for i := 0; i < rows+1; i++ {
		dp[i] = make([]int, cols+1)
	}
	dp[0][0] = 1
	for i := 1; i < cols+1; i++ {
		if p[i-1] == '*' {
			dp[0][i] = dp[0][i-2]
		} else {
			dp[0][i] = 0
		}
	}

	for i := 1; i < rows+1; i++ {
		for j := 1; j < cols+1; j++ {
			if p[j-1] == '*' {
				if p[j-2] == s[i-1] || p[j-2] == '.' {
					if dp[i][j-2] == 1 || dp[i-1][j] == 1 {
						dp[i][j] = 1
					} else {
						dp[i][j] = 0
					}
				} else {
					dp[i][j] = dp[i][j-2]
				}
			} else if p[j-1] == s[i-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 0
			}
		}

	}
	return dp[rows][cols] == 1

}

func main() {
	//s, p := "aa", "a"
	//fmt.Println(isMatch(s, p))
	//s, p = "aa", "a*"
	//fmt.Println(isMatch(s, p))
	//s, p = "ab", ".*"
	//fmt.Println(isMatch(s, p))
	//s, p = "mississippi", "mis*is*ip*."
	//fmt.Println(isMatch(s, p))
	s, p := "aab", "c*a*b"
	fmt.Println(isMatch(s, p))
	s, p = "ab", ".*c"
	fmt.Println(isMatch(s, p))

}

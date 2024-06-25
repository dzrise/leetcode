/*
387. First Unique Character in a String
Medium
Topics
Companies
Given a string s, find the first non-repeating character in it and return its index. If it does not exist, return -1.

Example 1:

Input: s = "leetcode"
Output: 0
Example 2:

Input: s = "loveleetcode"
Output: 2
Example 3:

Input: s = "aabb"
Output: -1

Constraints:

1 <= s.length <= 105
s consists of only lowercase English letters.
*/
package main

import "fmt"

func firstUniqChar(s string) int {
	if len(s) == 1 {
		return 0
	}

	var mS = make(map[uint8][2]int)
	for i, _ := range s {
		if _, ok := mS[s[i]]; !ok {
			mS[s[i]] = [2]int{i, 1}
		} else {
			newI := mS[s[i]][1] + 1
			mS[s[i]] = [2]int{i, newI}
		}
	}

	res := -1
	for _, v := range mS {
		if v[1] == 1 {
			if res == -1 || res > v[0] {
				res = v[0]
			}
		}
	}

	return res
}

func main() {
	fmt.Println(firstUniqChar("dddccdbba"))
}

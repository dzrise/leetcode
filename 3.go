// Example 1:
//
// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.
// Example 2:
//
// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.
// Example 3:
//
// Input: s = "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
//
// Constraints:
//
// 0 <= s.length <= 5 * 104
// s consists of English letters, digits, symbols and spaces.
package main

import "fmt"

func lengthOfLongestSubstringOld(s string) int {
	if len(s) == 0 {
		return 0
	}

	if len(s) == 1 {
		return 1
	}

	m := make(map[byte]bool)
	start := 0
	max := 0
	i := 0
	for {
		if i == len(s) {
			if i-start > max {
				max = i - start
			}
			break
		}

		if _, ok := m[s[i]]; ok {
			for k := range m {
				delete(m, k)
			}
			if i-start > max {
				max = i - start
			}
			start++
			i = start

		}

		m[s[i]] = true

		i++
	}

	return max
}

func lengthOfLongestSubstring(s string) int {
	max := 0
	start := 0
	m := map[byte]int{}
	end := 0
	for end = 0; end < len(s); end++ {
		m[s[end]]++
		for m[s[end]] > 1 {
			m[s[start]]--
			start++
		}
		if end-start+1 > max {
			max = end - start + 1
		}
		fmt.Printf("max: %d, end: %d, start: %d\n", max, end, start)
	}

	return max
}

func main() {
	s := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))
}

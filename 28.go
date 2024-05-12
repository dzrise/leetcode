// 28. Find the Index of the First Occurrence in a String
// Easy
// Topics
// Companies
// Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
//
// Example 1:
//
// Input: haystack = "sadbutsad", needle = "sad"
// Output: 0
// Explanation: "sad" occurs at index 0 and 6.
// The first occurrence is at index 0, so we return 0.
// Example 2:
//
// Input: haystack = "leetcode", needle = "leeto"
// Output: -1
// Explanation: "leeto" did not occur in "leetcode", so we return -1.
//
// Constraints:
//
// 1 <= haystack.length, needle.length <= 104
// haystack and needle consist of only lowercase English characters.
package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	if len(haystack) < len(needle) {
		return -1
	}

	if len(haystack) == len(needle) {
		if haystack == needle {
			return 0
		} else {
			return -1
		}
	}

	index := 0
	cur := 0
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[cur] {
			if cur == 0 {
				index = i
			}

			cur++
		} else {
			if cur != 0 {
				cur = 0
				i = index
			}
		}

		if cur == len(needle) {
			return index
		}

	}

	return -1
}

func main() {
	fmt.Println(strStr("leetcode", "leeto"))
	fmt.Println(strStr("sadbutsad", "sad"))
	fmt.Println(strStr("mississippi", "issip"))
	fmt.Println(strStr("mississippi", "pi"))
}

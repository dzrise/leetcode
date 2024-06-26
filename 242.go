/*
242. Valid Anagram
Easy
Topics
Companies
Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

Example 1:

Input: s = "anagram", t = "nagaram"
Output: true
Example 2:

Input: s = "rat", t = "car"
Output: false

Constraints:

1 <= s.length, t.length <= 5 * 104
s and t consist of lowercase English letters.

Follow up: What if the inputs contain Unicode characters? How would you adapt your solution to such a case?
*/
package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	if s == t {
		return true
	}

	summs := make(map[byte]int)
	summt := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		summs[s[i]]++
		summt[t[i]]++
	}

	for i := 0; i < len(s); i++ {
		if summs[s[i]] != summt[s[i]] {
			return false
		}
	}

	return true

}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
}

// 2423. Remove Letter To Equalize Frequency
// Easy
// Topics
// Companies
// Hint
// You are given a 0-indexed string word, consisting of lowercase English letters. You need to select one index and remove the letter at that index from word so that the frequency of every letter present in word is equal.
//
// Return true if it is possible to remove one letter so that the frequency of all letters in word are equal, and false otherwise.
//
// Note:
//
// The frequency of a letter x is the number of times it occurs in the string.
// You must remove exactly one letter and cannot choose to do nothing.
//
// Example 1:
//
// Input: word = "abcc"
// Output: true
// Explanation: Select index 3 and delete it: word becomes "abc" and each character has a frequency of 1.
// Example 2:
//
// Input: word = "aazz"
// Output: false
// Explanation: We must delete a character, so either the frequency of "a" is 1 and the frequency of "z" is 2, or vice versa. It is impossible to make all present letters have equal frequency.
//
// Constraints:
//
// 2 <= word.length <= 100
// word consists of lowercase English letters only.
package main

import (
	"fmt"
)

func equalFrequency(word string) bool {
	m := make(map[int32]int)
	for _, v := range word {
		m[v]++
	}

	maxOne := 1
	res := false
	for _, v := range m {
		if v == 1 || v-1%2 == 0 || v-1 == 1 {
			res = true
		} else {
			res = false
		}
		if v > maxOne {
			maxOne = v
		}
	}

	if maxOne == 1 {
		return true
	}

	return res
}

func main() {
	fmt.Println(equalFrequency("абкк"))
	fmt.Println(equalFrequency("баббдд"))
	fmt.Println(equalFrequency("aazz"))
	fmt.Println(equalFrequency("cccaa"))
}

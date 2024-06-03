// 49. Group Anagrams
// Medium
// Topics
// Companies
// Given an array of strings strs, group the anagrams together. You can return the answer in any order.
//
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.
//
// Example 1:
//
// Input: strs = ["eat","tea","tan","ate","nat","bat"]
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
// Example 2:
//
// Input: strs = [""]
// Output: [[""]]
// Example 3:
//
// Input: strs = ["a"]
// Output: [["a"]]
//
// Constraints:
//
// 1 <= strs.length <= 104
// 0 <= strs[i].length <= 100
// strs[i] consists of lowercase English letters.
package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {

	if len(strs) == 0 {
		return [][]string{}
	}

	groupsByCharSort := make(map[string][]string)

	for _, str := range strs {
		chars := []byte(str)
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})
		charsStr := string(chars)
		if _, ok := groupsByCharSort[charsStr]; !ok {
			groupsByCharSort[charsStr] = []string{str}
		} else {
			groupsByCharSort[charsStr] = append(groupsByCharSort[charsStr], str)
		}
	}

	res := make([][]string, 0, len(groupsByCharSort))

	for _, group := range groupsByCharSort {
		res = append(res, group)
	}

	return res

}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(groupAnagrams([]string{""}))
	fmt.Println(groupAnagrams([]string{"a"}))
}

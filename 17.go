// Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.
//
// A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.
//
// Example 1:
//
// Input: digits = "23"
// Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
// Example 2:
//
// Input: digits = ""
// Output: []
// Example 3:
//
// Input: digits = "2"
// Output: ["a","b","c"]
//
// Constraints:
//
// 0 <= digits.length <= 4
// digits[i] is a digit in the range ['2', '9'].
package main

import (
	"fmt"
	"strconv"
)

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	data := []string{"0", "1", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	res := []string{""}
	for i, _ := range digits {
		index, _ := strconv.Atoi(string(digits[i]))
		str := data[index]
		temp := []string{}
		for s := range res {
			for j, _ := range str {
				temp = append(temp, res[s]+string(str[j]))
			}
		}

		res = temp
	}

	return res
}

func main() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations(""))
	fmt.Println(letterCombinations("2"))
}

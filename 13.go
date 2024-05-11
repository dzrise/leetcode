// Code
// Testcase
// Testcase
// Test Result
// 13. Roman to Integer
// Easy
// Topics
// Companies
// Hint
// Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
//
// Symbol       Value
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000
// For example, 2 is written as II in Roman numeral, just two ones added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.
//
// Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:
//
// I can be placed before V (5) and X (10) to make 4 and 9.
// X can be placed before L (50) and C (100) to make 40 and 90.
// C can be placed before D (500) and M (1000) to make 400 and 900.
// Given a roman numeral, convert it to an integer.
//
// Example 1:
//
// Input: s = "III"
// Output: 3
// Explanation: III = 3.
// Example 2:
//
// Input: s = "LVIII"
// Output: 58
// Explanation: L = 50, V= 5, III = 3.
// Example 3:
//
// Input: s = "MCMXCIV"
// Output: 1994
// Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
//
// Constraints:
//
// 1 <= s.length <= 15
// s contains only the characters ('I', 'V', 'X', 'L', 'C', 'D', 'M').
// It is guaranteed that s is a valid roman numeral in the range [1, 3999].
package main

import "fmt"

func romanToInt(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		switch string(s[i]) {
		case "M":
			res += 1000
		case "D":
			res += 500
		case "C":
			if i != len(s)-1 && string(s[i+1]) == "M" {
				res += 900
				i++
			} else if i != len(s)-1 && string(s[i+1]) == "D" {
				res += 400
				i++
			} else {
				res += 100
			}
		case "L":
			res += 50
		case "X":
			if i != len(s)-1 && string(s[i+1]) == "C" {
				res += 90
				i++
			} else if i != len(s)-1 && string(s[i+1]) == "L" {
				res += 40
				i++
			} else {
				res += 10
			}
		case "V":
			res += 5
		case "I":
			if i != len(s)-1 && string(s[i+1]) == "X" {
				res += 9
				i++
			} else if i != len(s)-1 && string(s[i+1]) == "V" {
				res += 4
				i++
			} else {
				res += 1
			}

		}
	}

	return res
}

//func romanToInt(s string) int {
//	ret := 0
//	mp := map[byte]int{
//		'I': 1,
//		'V': 5,
//		'X': 10,
//		'L': 50,
//		'C': 100,
//		'D': 500,
//		'M': 1000,
//	}
//	mp2 := map[string]int{
//		"IV": 4,
//		"IX": 9,
//		"XL": 40,
//		"XC": 90,
//		"CD": 400,
//		"CM": 900,
//	}
//	for i := 0; i < len(s); i++ {
//		if i+1 < len(s) {
//			v, find := mp2[s[i:i+2]]
//			if find {
//				ret += v
//				i++
//				continue
//			}
//		}
//		ret += mp[s[i]]
//	}
//	return ret
//}

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
	fmt.Println(romanToInt("IX"))
}

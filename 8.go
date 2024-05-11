// Implement the myAtoi(string s) function, which converts a string to a 32-bit signed integer.
//
// The algorithm for myAtoi(string s) is as follows:
//
// Whitespace: Ignore any leading whitespace (" ").
// Signedness: Determine the sign by checking if the next character is '-' or '+', assuming positivity is neither present.
// Conversion: Read the integer by skipping leading zeros until a non-digit character is encountered or the end of the string is reached. If no digits were read, then the result is 0.
// Rounding: If the integer is out of the 32-bit signed integer range [-231, 231 - 1], then round the integer to remain in the range. Specifically, integers less than -231 should be rounded to -231, and integers greater than 231 - 1 should be rounded to 231 - 1.
// Return the integer as the final result.
//
// Example 1:
//
// Input: s = "42"
//
// Output: 42
//
// Explanation:
//
// The underlined characters are what is read in and the caret is the current reader position.
// Step 1: "42" (no characters read because there is no leading whitespace)
// ^
// Step 2: "42" (no characters read because there is neither a '-' nor '+')
// ^
// Step 3: "42" ("42" is read in)
// ^
// Example 2:
//
// Input: s = " -042"
//
// Output: -42
//
// Explanation:
//
// Step 1: "   -042" (leading whitespace is read and ignored)
// ^
// Step 2: "   -042" ('-' is read, so the result should be negative)
// ^
// Step 3: "   -042" ("042" is read in, leading zeros ignored in the result)
// ^
// Example 3:
//
// Input: s = "1337c0d3"
//
// Output: 1337
//
// Explanation:
//
// Step 1: "1337c0d3" (no characters read because there is no leading whitespace)
// ^
// Step 2: "1337c0d3" (no characters read because there is neither a '-' nor '+')
// ^
// Step 3: "1337c0d3" ("1337" is read in; reading stops because the next character is a non-digit)
// ^
// Example 4:
//
// Input: s = "0-1"
//
// Output: 0
//
// Explanation:
//
// Step 1: "0-1" (no characters read because there is no leading whitespace)
// ^
// Step 2: "0-1" (no characters read because there is neither a '-' nor '+')
// ^
// Step 3: "0-1" ("0" is read in; reading stops because the next character is a non-digit)
// ^
// Example 5:
//
// Input: s = "words and 987"
//
// Output: 0
//
// Explanation:
//
// Reading stops at the first non-digit character 'w'.
//
// Constraints:
//
// 0 <= s.length <= 200
// s consists of English letters (lower-case and upper-case), digits (0-9), ' ', '+', '-', and '.'.
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func myAtoi(s string) int {
	res := 0
	minusPlus := 1

	s = strings.TrimSpace(s)

	if len(s) == 0 {
		return 0
	}

	if s[:1] == "-" {
		minusPlus = -1
		s = s[1:]
	} else if s[:1] == "+" {
		s = s[1:]
	}

	for i := 0; i < len(s); i++ {
		if !unicode.IsDigit(rune(s[i])) {
			break
		}
		res = res*10 + int(s[i]-'0')

		if res*minusPlus < -2147483648 {
			return -2147483648
		}

		if res*minusPlus > 2147483647 {
			return 2147483647
		}

	}

	return minusPlus * res
}

//func myAtoi(s string) int {
//	c := map[string]bool{"0": true, "1": true, "2": true, "3": true, "4": true, "5": true, "6": true, "7": true, "8": true, "9": true}
//	res := 0
//	minusPlus := 1
//
//	s = strings.TrimSpace(s)
//
//	if len(s) == 0 {
//		return 0
//	}
//
//	if s[:1] == "-" {
//		minusPlus = -1
//		s = s[1:]
//	} else if s[:1] == "+" {
//		s = s[1:]
//	}
//
//	for i := 0; i < len(s); i++ {
//		if c[string(s[i])] {
//			res = res*10 + int(s[i]-'0')
//		} else {
//			break
//		}
//
//		if res*minusPlus < -2147483648 {
//			return -2147483648
//		}
//
//		if res*minusPlus > 2147483647 {
//			return 2147483647
//		}
//
//	}
//
//	return minusPlus * res
//}

//
//
//func myAtoi(s string) int {
//	c := map[string]bool{"0": true, "1": true, "2": true, "3": true, "4": true, "5": true, "6": true, "7": true, "8": true, "9": true}
//	res := ""
//	needMinus := false
//	stop := false
//	for i := 0; i < len(s); i++ {
//		if string(s[i]) == " " && res == "" {
//			if stop {
//				break
//			}
//			continue
//		}
//		if string(s[i]) == "+" || string(s[i]) == "-" {
//			if res != "" || stop {
//				break
//			} else {
//				if string(s[i]) == "-" {
//					needMinus = true
//				}
//				res += string(s[i])
//				continue
//			}
//		}
//
//		if c[string(s[i])] {
//			if string(s[i]) == "0" {
//				if res == "" || res == "-" || res == "+" {
//					stop = true
//					continue
//				}
//			}
//			res += string(s[i])
//		} else {
//
//			if string(s[i]) != "-" || res != "" || stop {
//				break
//			}
//
//		}
//
//	}
//
//	if len(res) == 0 {
//		return 0
//	}
//	if res[0:1] == "+" {
//		res = res[1:]
//	}
//
//	if len(res) > 11 && needMinus {
//		return -2147483648
//	}
//
//	if len(res) > 10 && !needMinus {
//		return 2147483647
//	}
//
//	rInt := 0
//	_, err := fmt.Sscan(res, &rInt)
//	if err != nil {
//		return 0
//	}
//
//	if rInt > 2147483647 {
//		return 2147483647
//	}
//
//	if rInt < -2147483648 {
//		return -2147483648
//	}
//
//	return rInt
//}

func main() {
	//fmt.Println(myAtoi("-+12"))
	//fmt.Println(myAtoi("-+12"))
	fmt.Println(myAtoi("    +1146905820n1"))
	//fmt.Println(myAtoi("0-1"))
	//fmt.Println(myAtoi("0-1"))
	//fmt.Println(myAtoi("words and 987"))
}

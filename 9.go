// 9. Palindrome Number
// Easy
// Topics
// Companies
// Hint
// Given an integer x, return true if x is a
// palindrome
// , and false otherwise.
//
// Example 1:
//
// Input: x = 121
// Output: true
// Explanation: 121 reads as 121 from left to right and from right to left.
// Example 2:
//
// Input: x = -121
// Output: false
// Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
// Example 3:
//
// Input: x = 10
// Output: false
// Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
//
// Constraints:
//
// -231 <= x <= 231 - 1
//
// Follow up: Could you solve it without converting the integer to a string?
package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	if x%10 == 0 {
		return false
	}

	part := 0
	for x > 0 {
		part = part*10 + x%10

		x /= 10
		fmt.Println(x)
		fmt.Println(part)
		if part == x || part == x/10 {
			return true
		}
	}

	return false
}

//func isPalindrome(x int) bool {
//	xStr := strconv.Itoa(x)
//	if x < 0 {
//		return false
//	}
//	left := 0
//	right := len(xStr) - 1
//	for left < right {
//		if xStr[left] != xStr[right] {
//			return false
//		}
//		left++
//		right--
//	}
//	return true
//}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
	fmt.Println(isPalindrome(111344))
}

//125. Valid Palindrome
//Easy
//Topics
//Companies
//A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.
//
//Given a string s, return true if it is a palindrome, or false otherwise.
//
//
//
//Example 1:
//
//Input: s = "A man, a plan, a canal: Panama"
//Output: true
//Explanation: "amanaplanacanalpanama" is a palindrome.
//Example 2:
//
//Input: s = "race a car"
//Output: false
//Explanation: "raceacar" is not a palindrome.
//Example 3:
//
//Input: s = " "
//Output: true
//Explanation: s is an empty string "" after removing non-alphanumeric characters.
//Since an empty string reads the same forward and backward, it is a palindrome.
//
//
//Constraints:
//
//1 <= s.length <= 2 * 105
//s consists only of printable ASCII characters.

package main

import (
	"fmt"
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	var re = regexp.MustCompile(`[[:punct:]]|[[:space:]]`)
	s = re.ReplaceAllString(s, "")
	s = strings.ToLower(s)

	if len(s) <= 1 {
		return true
	}

	j := len(s) - 1
	i := 0
	for i < j {
		if string(s[i]) == " " {
			i++
			continue
		}
		if string(s[j]) == " " {
			j--
			continue
		}

		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome(" "))
}

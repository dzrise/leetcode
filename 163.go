/*
168. Excel Sheet Column Title
Easy
Topics
Companies
Given an integer columnNumber, return its corresponding column title as it appears in an Excel sheet.

For example:

A -> 1
B -> 2
C -> 3
...
Z -> 26
AA -> 27
AB -> 28
...

Example 1:

Input: columnNumber = 1
Output: "A"
Example 2:

Input: columnNumber = 28
Output: "AB"
Example 3:

Input: columnNumber = 701
Output: "ZY"

Constraints:

1 <= columnNumber <= 231 - 1
*/
package main

import (
	"fmt"
)

func convertToTitle(columnNumber int) string {
	alf := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	if columnNumber < len(alf) {
		return alf[columnNumber-1]
	}
	res := []string{}
	for columnNumber > 0 {
		res = append(res, alf[(columnNumber-1)%26])
		columnNumber = (columnNumber - 1) / 26
	}

	res_str := ""
	for i := len(res); i > 0; i-- {
		res_str += string(res[i-1])
	}

	return res_str
}

func main() {
	//fmt.Println(convertToTitle(28))
	fmt.Println(convertToTitle(701))

	fmt.Println(convertToTitle(2147483647))
}

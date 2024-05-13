// 6. Zigzag Conversion
// Solved
// Medium
// Topics
// Companies
// The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)
//
// P   A   H   N
// A P L S I I G
// Y   I   R
// And then read line by line: "PAHNAPLSIIGYIR"
//
// Write the code that will take a string and make this conversion given a number of rows:
//
// string convert(string s, int numRows);
//
// Example 1:
//
// Input: s = "PAYPALISHIRING", numRows = 3
// Output: "PAHNAPLSIIGYIR"
// Example 2:
//
// Input: s = "PAYPALISHIRING", numRows = 4
// Output: "PINALSIGYAHRPI"
// Explanation:
// P     I    N
// A   L S  I G
// Y A   H R
// P     I
// Example 3:
//
// Input: s = "A", numRows = 1
// Output: "A"
//
// Constraints:
//
// 1 <= s.length <= 1000
// s consists of English letters (lower-case and upper-case), ',' and '.'.
// 1 <= numRows <= 1000
package main

import "fmt"

func convert(s string, numRows int) string {
	if len(s) <= numRows || numRows == 1 {
		return s
	}

	res := make([][]byte, numRows)
	flag := 1
	row := 0

	for i := 0; i < len(s); i++ {
		res[row] = append(res[row], s[i])
		if row == 0 {
			flag = 1
		} else if row == numRows-1 {
			flag = -1
		}
		row = row + flag
	}

	resStr := ""

	for i := 0; i < numRows; i++ {
		resStr += string(res[i])
	}

	return resStr

}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert("PAYPALISHIRING", 4))
	fmt.Println(convert("A", 1))
}

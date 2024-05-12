//67. Add Binary
//Easy
//Topics
//Companies
//Given two binary strings a and b, return their sum as a binary string.
//
//
//
//Example 1:
//
//Input: a = "11", b = "1"
//Output: "100"
//Example 2:
//
//Input: a = "1010", b = "1011"
//Output: "10101"
//
//
//Constraints:
//
//1 <= a.length, b.length <= 104
//a and b consist only of '0' or '1' characters.
//Each string does not contain leading zeros except for the zero itself.

package main

import "fmt"

func addBinary(a string, b string) string {
	if len(a) == 0 {
		return b
	}

	if len(b) == 0 {
		return a
	}

}

func main() {
	fmt.Println(addBinary("11", "1"))
	fmt.Println(addBinary("1010", "1011"))
}

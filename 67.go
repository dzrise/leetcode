// 67.Given two binary strings a and b, return their sum as a binary string.
//
// Example 1:
//
// Input: a = "11", b = "1"
// Output: "100"
// Example 2:
//
// Input: a = "1010", b = "1011"
// Output: "10101"
//
// Constraints:
//
// 1 <= a.length, b.length <= 104
// a and b consist only of '0' or '1' characters.
// Each string does not contain leading zeros except for the zero itself.
package main

import (
	"fmt"
	"math/big"
)

func addBinary(a string, b string) string {
	num1, _ := new(big.Int).SetString(a, 2)
	num2, _ := new(big.Int).SetString(b, 2)
	num1.Add(num1, num2)
	return num1.Text(2)
}

//func Reverse(s string) string {
//	runes := []rune(s)
//	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
//		runes[i], runes[j] = runes[j], runes[i]
//	}
//	return string(runes)
//}
//
//func addBinary(a string, b string) string {
//	if a == "" {
//		return b
//	}
//
//	if b == "" {
//		return a
//	}
//	bg := ""
//	mn := ""
//	if len(a) < len(b) {
//		bg = Reverse(b)
//		mn = Reverse(a)
//	} else {
//		bg = Reverse(a)
//		mn = Reverse(b)
//	}
//	res := ""
//	carry := '0'
//	for i := 0; i < len(bg); i++ {
//		m := '0'
//		if i < len(mn) {
//			m = int32(mn[i])
//		}
//		if bg[i] == '1' && m == '1' {
//			if carry == '0' {
//				res += "0"
//			} else {
//				res += "1"
//			}
//			carry = '1'
//		} else if bg[i] == '0' && m == '0' {
//			if carry == '1' {
//				res += "1"
//			} else {
//				res += "0"
//			}
//			carry = '0'
//		} else if bg[i] == '0' || m == '0' {
//			if carry == '1' {
//				res += "0"
//				carry = '1'
//			} else {
//				res += "1"
//			}
//		}
//
//		if i == len(bg)-1 && carry == '1' {
//			res += "1"
//		}
//
//	}
//	return Reverse(res)
//}

func main() {
	fmt.Println(addBinary("11", "1"))
	fmt.Println(addBinary("1010", "1011"))
}

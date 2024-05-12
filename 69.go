// 69. Sqrt(x)
// Easy
// Topics
// Companies
// Hint
// Given a non-negative integer x, return the square root of x rounded down to the nearest integer. The returned integer should be non-negative as well.
//
// You must not use any built-in exponent function or operator.
//
// For example, do not use pow(x, 0.5) in c++ or x ** 0.5 in python.
//
// Example 1:
//
// Input: x = 4
// Output: 2
// Explanation: The square root of 4 is 2, so we return 2.
// Example 2:
//
// Input: x = 8
// Output: 2
// Explanation: The square root of 8 is 2.82842..., and since we round it down to the nearest integer, 2 is returned.
//
// Constraints:
//
// 0 <= x <= 231 - 1
package main

import (
	"fmt"
)

func mySqrt(x int) int {
	left, right := 0, x
	for right-left > 1 {
		mid := left + (right-left)/2
		if mid*mid <= x {
			left = mid
		} else if mid*mid > x {
			right = mid
		}
	}

	return left
}

//func mySqrt(x int) int {
//	i := 1
//	for {
//		if i*i > x {
//			return i - 1
//		}
//		i++
//	}
//}

func main() {
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(8))
}

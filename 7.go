//Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.
//
//Assume the environment does not allow you to store 64-bit integers (signed or unsigned).
//
//
//
//Example 1:
//
//Input: x = 123
//Output: 321
//Example 2:
//
//Input: x = -123
//Output: -321
//Example 3:
//
//Input: x = 120
//Output: 21

package main

import (
	"fmt"
	"strconv"
)

func reverse(x int) int {
	res := 0
	for x != 0 {
		res = res*10 + x%10
		x = x / 10
	}
	if res > 2147483647 || res < -2147483648 {
		return 0
	}
	return res
}

func reverseOld(x int) int {
	str := fmt.Sprintf("%d", x)

	res := ""

	for _, v := range str {
		if string(v) == "-" {
			continue
		}
		res = string(v) + res
	}

	if x < 0 {
		res = "-" + res
	}

	resInt, _ := strconv.Atoi(res)
	if resInt > 2147483647 || resInt < -2147483648 {
		return 0
	}

	return resInt
}

func main() {
	fmt.Println(reverseOld(-123))
	fmt.Println(reverseOld(120))
}

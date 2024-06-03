/*
179. Largest Number
Medium
Topics
Companies
Given a list of non-negative integers nums, arrange them such that they form the largest number and return it.

Since the result may be very large, so you need to return a string instead of an integer.

Example 1:

Input: nums = [10,2]
Output: "210"
Example 2:

Input: nums = [3,30,34,5,9]
Output: "9534330"

Constraints:

1 <= nums.length <= 100
0 <= nums[i] <= 109
*/
package main

import (
	"fmt"
	"sort"
	"strings"
)

func largestNumber(nums []int) string {
	if len(nums) == 0 {
		return ""
	}
	str := ""
	strs := []string{}
	for _, v := range nums {
		strs = append(strs, fmt.Sprintf("%d", v))
	}

	sort.Slice(strs, func(i, j int) bool {
		if strs[i][0] < strs[j][0] {
			return false
		}

		if strs[i][0] == strs[j][0] {
			if strs[i]+strs[j] < strs[j]+strs[i] {
				return false
			}

		}

		return true
	})
	for _, s := range strs {
		if str == "0" && s == "0" {
			continue
		}
		str += strings.TrimSpace(s)
	}
	return str
}

func main() {
	fmt.Println(largestNumber([]int{10, 2}))
	fmt.Println(largestNumber([]int{3, 30, 34, 5, 9}))
}

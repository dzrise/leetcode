/*
594. Longest Harmonious Subsequence
Easy
Topics
Companies
We define a harmonious array as an array where the difference between its maximum value and its minimum value is exactly 1.

Given an integer array nums, return the length of its longest harmonious subsequence among all its possible subsequences.

A subsequence of array is a sequence that can be derived from the array by deleting some or no elements without changing the order of the remaining elements.

Example 1:

Input: nums = [1,3,2,2,5,2,3,7]
Output: 5
Explanation: The longest harmonious subsequence is [3,2,2,2,3].
Example 2:

Input: nums = [1,2,3,4]
Output: 2
Example 3:

Input: nums = [1,1,1,1]
Output: 0

Constraints:

1 <= nums.length <= 2 * 104
-109 <= nums[i] <= 109
*/
package main

import (
	"fmt"
	"sort"
)

func findLHS(nums []int) int {
	m, c := 0, 1
	sort.Ints(nums)
	minIndex := 0
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[minIndex] == 1 || nums[i]-nums[minIndex] == -1 || nums[i]-nums[minIndex] == 0 {
			c++
			if c > m {
				m = c
			}
		} else {
			minIndex += 1
			i = minIndex + 2
			c = 1
		}
	}

	return m
}

func main() {
	fmt.Println(findLHS([]int{1, 3, 2, 2, 5, 2, 3, 7}))
}

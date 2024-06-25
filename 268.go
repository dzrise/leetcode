/*
268. Missing Number
Medium
Topics
Companies
Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array.

Example 1:

Input: nums = [3,0,1]
Output: 2
Explanation: n = 3 since there are 3 numbers, so all numbers are in the range [0,3]. 2 is the missing number in the range since it does not appear in nums.
Example 2:

Input: nums = [0,1]
Output: 2
Explanation: n = 2 since there are 2 numbers, so all numbers are in the range [0,2]. 2 is the missing number in the range since it does not appear in nums.
Example 3:

Input: nums = [9,6,4,2,3,5,7,0,1]
Output: 8
Explanation: n = 9 since there are 9 numbers, so all numbers are in the range [0,9]. 8 is the missing number in the range since it does not appear in nums.

Constraints:

n == nums.length
1 <= n <= 104
*/
package main

import (
	"fmt"
	"sort"
)

func missingNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		if nums[0] == 1 {
			return 0
		}
		return nums[0] + 1
	}

	sort.Ints(nums)

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1]+1 {
			return nums[i] - 1
		}

		if i == len(nums)-1 {
			if nums[0] == 1 {
				return 0
			}
			return nums[i] + 1
		}
	}
	return 0
}

func main() {
	fmt.Println(missingNumber([]int{3, 0, 1}))
}

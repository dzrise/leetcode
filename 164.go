/*
164. Maximum Gap
Medium
Topics
Companies
Given an integer array nums, return the maximum difference between two successive elements in its sorted form. If the array contains less than two elements, return 0.

You must write an algorithm that runs in linear time and uses linear extra space.

Example 1:

Input: nums = [3,6,9,1]
Output: 3
Explanation: The sorted form of the array is [1,3,6,9], either (3,6) or (6,9) has the maximum difference 3.
Example 2:

Input: nums = [10]
Output: 0
Explanation: The array contains less than 2 elements, therefore return 0.

Constraints:

1 <= nums.length <= 105
0 <= nums[i] <= 109
*/
package main

import (
	"fmt"
	"sort"
)

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	maxGap := 0

	sort.Ints(nums)

	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] > maxGap {
			maxGap = nums[i] - nums[i-1]
		}
	}

	return maxGap

}

func main() {
	fmt.Println(maximumGap([]int{3, 6, 9, 1}))
	fmt.Println(maximumGap([]int{10}))
}

/*
Code
Testcase
Testcase
Test Result
75. Sort Colors
Medium
Topics
Companies
Hint
Given an array nums with n objects colored red, white, or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white, and blue.

We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.

You must solve this problem without using the library's sort function.

Example 1:

Input: nums = [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]
Example 2:

Input: nums = [2,0,1]
Output: [0,1,2]

Constraints:

n == nums.length
1 <= n <= 300
nums[i] is either 0, 1, or 2.

Follow up: Could you come up with a one-pass algorithm using only constant extra space?
*/
package main

import "fmt"

func sortColors(nums []int) {
	if len(nums) == 0 {
		return
	}
	i, j, k := 0, 0, len(nums)-1
	for j <= k {
		if nums[j] < 1 {
			nums[i], nums[j] = nums[j], nums[i]
			i, j = i+1, j+1
		} else if nums[j] > 1 {
			nums[j], nums[k] = nums[k], nums[j]
			k--
		} else {
			j++
		}
	}
}

func main() {
	arr1 := []int{2, 0, 2, 1, 1, 0}
	sortColors(arr1)
	fmt.Println(arr1)
	arr2 := []int{2, 0, 1}
	sortColors(arr2)
	fmt.Println(arr2)
}

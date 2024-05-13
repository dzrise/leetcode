// 33. Search in Rotated Sorted Array
// Medium
// Topics
// Companies
// There is an integer array nums sorted in ascending order (with distinct values).
//
// Prior to being passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].
//
// Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums, or -1 if it is not in nums.
//
// You must write an algorithm with O(log n) runtime complexity.
//
// Example 1:
//
// Input: nums = [4,5,6,7,0,1,2], target = 0
// Output: 4
// Example 2:
//
// Input: nums = [4,5,6,7,0,1,2], target = 3
// Output: -1
// Example 3:
//
// Input: nums = [1], target = 0
// Output: -1
//
// Constraints:
//
// 1 <= nums.length <= 5000
// -104 <= nums[i] <= 104
// All values of nums are unique.
// nums is an ascending array that is possibly rotated.
// -104 <= target <= 104
package main

import "fmt"

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	if len(nums) == 1 {
		if target == nums[0] {
			return 0
		} else {
			return -1
		}
	}
	if len(nums) == 2 {
		if target == nums[0] {
			return 0
		} else if target == nums[1] {
			return 1
		} else {
			return -1
		}
	}

	left := len(nums) / 2
	for left > 0 {
		for i, n := range nums[left:] {
			if n == target {
				return i + left
			}
		}
		if len(nums[:left+1]) > 2 {
			left = len(nums[:left+1]) / 2
		} else {
			if target == nums[0] {
				return 0
			} else if target == nums[1] {
				return 1
			} else {
				return -1
			}
		}

	}
	return -1
}

//func search(nums []int, target int) int {
//	for i, v := range nums {
//		if v == target {
//			return i
//		}
//	}
//
//	return -1
//}

func main() {
	fmt.Println(search([]int{3, 5, 1}, 1))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1}, 0))
	fmt.Println(search([]int{1, 3, 5}, 1))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1}, 3))
	fmt.Println(search([]int{1}, 0))
}

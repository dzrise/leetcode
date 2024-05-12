// 35. Search Insert Position
// Easy
// Topics
// Companies
// Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
//
// You must write an algorithm with O(log n) runtime complexity.
//
// Example 1:
//
// Input: nums = [1,3,5,6], target = 5
// Output: 2
// Example 2:
//
// Input: nums = [1,3,5,6], target = 2
// Output: 1
// Example 3:
//
// Input: nums = [1,3,5,6], target = 7
// Output: 4
//
// Constraints:
//
// 1 <= nums.length <= 104
// -104 <= nums[i] <= 104
// nums contains distinct values sorted in ascending order.
// -104 <= target <= 104
package main

import "fmt"

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		if nums[0] >= target {
			return 0
		} else {
			return 1
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
	}

	return len(nums)
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 5

	fmt.Println(searchInsert(nums, target))

	nums = []int{1, 3, 5, 6}
	target = 2

	fmt.Println(searchInsert(nums, target))

	nums = []int{1, 3, 5, 6}
	target = 7

	fmt.Println(searchInsert(nums, target))
}

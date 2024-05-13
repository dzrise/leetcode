// 34. Find First and Last Position of Element in Sorted Array
// Medium
// Topics
// Companies
// Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.
//
// If target is not found in the array, return [-1, -1].
//
// You must write an algorithm with O(log n) runtime complexity.
//
// Example 1:
//
// Input: nums = [5,7,7,8,8,10], target = 8
// Output: [3,4]
// Example 2:
//
// Input: nums = [5,7,7,8,8,10], target = 6
// Output: [-1,-1]
// Example 3:
//
// Input: nums = [], target = 0
// Output: [-1,-1]
//
// Constraints:
//
// 0 <= nums.length <= 105
// -109 <= nums[i] <= 109
// nums is a non-decreasing array.
// -109 <= target <= 109
package main

import "fmt"

func binSearchRange(nums []int, target int) []int {
	start := 0
	end := len(nums) - 1
	res := []int{-1, -1}

	if len(nums) == 0 {
		return res
	}

	if len(nums) == 1 {
		if nums[0] == target {
			return []int{0, 0}
		}
	}

	for end >= start {
		if nums[end] == target {
			if end > res[1] {
				res[1] = end
			}
			if end < res[0] || res[0] == -1 {
				res[0] = end
			}
		}

		if nums[start] == target {
			if start < res[0] || res[0] == -1 {
				res[0] = start
			}
			if start > res[1] || res[1] == -1 {
				res[1] = start
			}
		}
		fmt.Println(res)
		start++
		end--
	}

	return res
}
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	if len(nums) == 1 {
		if nums[0] == target {
			return []int{0, 0}
		} else {
			return []int{-1, -1}
		}
	}

	mid := len(nums) / 2
	res := []int{-1, -1}
	if nums[mid] < target {
		res = binSearchRange(nums[mid:], target)
		if res[0] != -1 {
			res[0] += mid
		}
		if res[1] != -1 {
			res[1] += mid
		}
		return res
	} else if nums[mid] > target {
		return binSearchRange(nums[:mid], target)
	} else {
		return binSearchRange(nums, target)
	}

}

func main() {
	//fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	//fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchRange([]int{1, 1, 3}, 1))
	//fmt.Println(searchRange([]int{}, 0))
}

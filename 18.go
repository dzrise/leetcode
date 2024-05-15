// 18. 4Sum
// Medium
// Topics
// Companies
// Given an array nums of n integers, return an array of all the unique quadruplets [nums[a], nums[b], nums[c], nums[d]] such that:
//
// 0 <= a, b, c, d < n
// a, b, c, and d are distinct.
// nums[a] + nums[b] + nums[c] + nums[d] == target
// You may return the answer in any order.
//
// Example 1:
//
// Input: nums = [1,0,-1,0,-2,2], target = 0
// Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
// Example 2:
//
// Input: nums = [2,2,2,2,2], target = 8
// Output: [[2,2,2,2]]
//
// Constraints:
//
// 1 <= nums.length <= 200
// -109 <= nums[i] <= 109
// -109 <= target <= 109
package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	resMap := make(map[[4]int][]int)
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-3; i++ {
		for j := i + 1; j < n-2; j++ {
			l := j + 1
			m := n - 1
			for l < m {
				if nums[i]+nums[j]+nums[l]+nums[m] == target {
					resMap[[4]int{nums[i], nums[j], nums[l], nums[m]}] = []int{nums[i], nums[j], nums[l], nums[m]}
					l++
					m--
				} else if nums[i]+nums[j]+nums[l]+nums[m] < target {
					l++
				} else {
					m--
				}
			}
		}
	}
	var res [][]int
	for _, v := range resMap {
		res = append(res, v)
	}

	return res
}

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))
	//fmt.Println(fourSum([]int{1,0,-1,0,-2,2}, 0))
}

// 16. 3Sum Closest
// Solved
// Medium
// Topics
// Companies
// Given an integer array nums of length n and an integer target, find three integers in nums such that the sum is closest to target.
//
// Return the sum of the three integers.
//
// You may assume that each input would have exactly one solution.
//
// Example 1:
//
// Input: nums = [-1,2,1,-4], target = 1
// Output: 2
// Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
// Example 2:
//
// Input: nums = [0,0,0], target = 1
// Output: 0
// Explanation: The sum that is closest to the target is 0. (0 + 0 + 0 = 0).
//
// Constraints:
//
// 3 <= nums.length <= 500
// -1000 <= nums[i] <= 1000
// -104 <= target <= 104
package main

import (
	"fmt"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}

	sort.Ints(nums)
	res := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				cur := nums[i] + nums[j] + nums[k]
				if cur == target {
					return target
				}
				if i == 0 && j == i+1 && k == j+1 {
					res = cur
				} else {
					a := target - cur
					b := target - res

					if a < 0 {
						a = -1 * a
					}
					if b < 0 {
						b = -1 * b
					}
					if a < b {
						res = cur
					}

				}

			}

		}
	}

	return res
}

func main() {
	//fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
	//fmt.Println(threeSumClosest([]int{4, 0, 5, -5, 3, 3, 0, -4, -5}, -2))
	//fmt.Println(threeSumClosest([]int{0, 1, 2}, 3))
	fmt.Println(threeSumClosest([]int{833, 736, 953, -584, -448, 207, 128, -445, 126, 248, 871, 860, 333, -899, 463, 488, -50, -331, 903, 575, 265, 162, -733, 648, 678, 549, 579, -172, -897, 562, -503, -508, 858, 259, -347, -162, -505, -694, 300, -40, -147, 383, -221, -28, -699, 36, -229, 960, 317, -585, 879, 406, 2, 409, -393, -934, 67, 71, -312, 787, 161, 514, 865, 60, 555, 843, -725, -966, -352, 862, 821, 803, -835, -635, 476, -704, -78, 393, 212, 767, -833, 543, 923, -993, 274, -839, 389, 447, 741, 999, -87, 599, -349, -515, -553, -14, -421, -294, -204, -713, 497, 168, 337, -345, -948, 145, 625, 901, 34, -306, -546, -536, 332, -467, -729, 229, -170, -915, 407, 450, 159, -385, 163, -420, 58, 869, 308, -494, 367, -33, 205, -823, -869, 478, -238, -375, 352, 113, -741, -970, -990, 802, -173, -977, 464, -801, -408, -77, 694, -58, -796, -599, -918, 643, -651, -555, 864, -274, 534, 211, -910, 815, -102, 24, -461, -146}, -7111))
	//fmt.Println(threeSumClosest([]int{1, 1, 1, 0}, 100))
}

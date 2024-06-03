/*
414. Third Maximum Number
Easy
Topics
Companies
Given an integer array nums, return the third distinct maximum number in this array. If the third maximum does not exist, return the maximum number.

Example 1:

Input: nums = [3,2,1]
Output: 1
Explanation:
The first distinct maximum is 3.
The second distinct maximum is 2.
The third distinct maximum is 1.
Example 2:

Input: nums = [1,2]
Output: 2
Explanation:
The first distinct maximum is 2.
The second distinct maximum is 1.
The third distinct maximum does not exist, so the maximum (2) is returned instead.
Example 3:

Input: nums = [2,2,3,1]
Output: 1
Explanation:
The first distinct maximum is 3.
The second distinct maximum is 2 (both 2's are counted together since they have the same value).
The third distinct maximum is 1.

Constraints:

1 <= nums.length <= 104
-231 <= nums[i] <= 231 - 1

Follow up: Can you find an O(n) solution?
*/
package main

import "fmt"

func thirdMax(nums []int) int {
	var res = make(map[int]int)
	for _, v := range nums {
		if n1, ok := res[2]; ok {
			if n1 == v {
				continue
			} else if v > n1 {
				res[2] = v
				v = n1
			}
		} else {
			res[2] = v
			continue

		}

		if n2, ok := res[1]; ok {
			if n2 == v {
				continue
			} else if v > n2 {
				res[1] = v
				v = n2
			}

		} else {
			res[1] = v
			continue
		}

		if n3, ok := res[0]; ok {
			if n3 == v {
				continue
			} else if v > n3 {
				res[0] = v
				continue
			}
		} else {
			res[0] = v
			continue
		}
		if v > res[2] {
			res[0] = res[1]
			res[1] = res[2]
			res[2] = v
		} else if v > res[1] {
			res[0] = res[1]
			res[1] = v
		} else if v > res[0] {
			res[0] = v
		}
	}
	if len(res) < 3 {
		return res[2]
	}

	return res[0]
}

func main() {
	fmt.Println(thirdMax([]int{3, 2, 1}))
	fmt.Println(thirdMax([]int{5, 2, 4, 1, 3, 6, 0}))
}

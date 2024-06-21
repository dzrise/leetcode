//228. Summary Ranges
//Easy
//Topics
//Companies
//You are given a sorted unique integer array nums.
//
//A range [a,b] is the set of all integers from a to b (inclusive).
//
//Return the smallest sorted list of ranges that cover all the numbers in the array exactly. That is, each element of nums is covered by exactly one of the ranges, and there is no integer x such that x is in one of the ranges but not in nums.
//
//Each range [a,b] in the list should be output as:
//
//"a->b" if a != b
//"a" if a == b
//
//
//Example 1:
//
//Input: nums = [0,1,2,4,5,7]
//Output: ["0->2","4->5","7"]
//Explanation: The ranges are:
//[0,2] --> "0->2"
//[4,5] --> "4->5"
//[7,7] --> "7"
//Example 2:
//
//Input: nums = [0,2,3,4,6,8,9]
//Output: ["0","2->4","6","8->9"]
//Explanation: The ranges are:
//[0,0] --> "0"
//[2,4] --> "2->4"
//[6,6] --> "6"
//[8,9] --> "8->9"
//
//
//Constraints:
//
//0 <= nums.length <= 20
//-231 <= nums[i] <= 231 - 1
//All the values of nums are unique.
//nums is sorted in ascending order.

package main

import (
	"fmt"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	if len(nums) == 1 {
		return []string{fmt.Sprintf("%d", nums[0])}
	}

	res := []string{}
	start := 0
	end := 0

	for i := 1; i < len(nums); i++ {
		if nums[i]-1 == nums[end] {
			end = i
			if len(nums)-1 == i {
				if nums[start] != nums[i]-1 {
					res = append(res, fmt.Sprintf("%d", nums[start]))
					res = append(res, fmt.Sprintf("%d", nums[i]))
				} else {
					res = append(res, fmt.Sprintf("%d->%d", nums[start], nums[i]))
				}
			}
		} else {
			if start < end {
				res = append(res, fmt.Sprintf("%d->%d", nums[start], nums[end]))
				if len(nums)-1 == i {
					res = append(res, fmt.Sprintf("%d", nums[i]))

				}
			} else if len(nums)-1 != i {
				res = append(res, fmt.Sprintf("%d", nums[start]))

			} else {
				if nums[start] != nums[i]-1 {
					res = append(res, fmt.Sprintf("%d", nums[start]))
					res = append(res, fmt.Sprintf("%d", nums[i]))
				} else {
					res = append(res, fmt.Sprintf("%d->%d", nums[start], nums[i]))
				}
			}
			start = i
			end = i
		}
	}

	//for i := 1; i < len(nums); i++ {
	//	n := nums[i]
	//	if n-1 != nums[i-1] {
	//		if start < i-1 {
	//			res = append(res, fmt.Sprintf("%d->%d", nums[start], nums[i-1]))
	//		} else if i != len(nums)-1 {
	//			res = append(res, fmt.Sprintf("%d", nums[start]))
	//		} else {
	//			res = append(res, fmt.Sprintf("%d->%d", nums[start], nums[i]))
	//		}
	//		start = i
	//	}
	//}

	return res
}

func main() {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
	fmt.Println(summaryRanges([]int{0, 1, 2}))
	//fmt.Println(summaryRanges([]int{0, 1, 2, 3}))
	//fmt.Println(summaryRanges([]int{0, 1, 2, 3, 4}))
	//fmt.Println(summaryRanges([]int{0, 1, 2, 3, 4, 5}))
}

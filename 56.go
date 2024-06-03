// 56. Merge Intervals
// Medium
// Topics
// Companies
// Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.
//
// Example 1:
//
// Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
// Output: [[1,6],[8,10],[15,18]]
// Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].
// Example 2:
//
// Input: intervals = [[1,4],[4,5]]
// Output: [[1,5]]
// Explanation: Intervals [1,4] and [4,5] are considered overlapping.
//
// Constraints:
//
// 1 <= intervals.length <= 104
// intervals[i].length == 2
// 0 <= starti <= endi <= 104
package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	if len(intervals) == 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{}

	cur := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= cur[1] {
			if intervals[i][0] < cur[0] {
				cur[0] = intervals[i][0]
			}

			if intervals[i][1] > cur[1] {
				cur[1] = intervals[i][1]
			}

			if i == len(intervals)-1 {
				res = append(res, cur)
			}
		} else {
			res = append(res, cur)
			cur = intervals[i]
			if i == len(intervals)-1 {
				res = append(res, intervals[i])
			}
		}

	}

	return res
}

func main() {
	fmt.Println(merge([][]int{[]int{1, 3}, []int{2, 6}, []int{8, 10}, []int{15, 18}}))
	fmt.Println(merge([][]int{[]int{1, 4}, []int{4, 5}}))
}

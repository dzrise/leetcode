/*
119. Pascal's Triangle II
Easy
Topics
Companies
Given an integer rowIndex, return the rowIndexth (0-indexed) row of the Pascal's triangle.

In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:

Example 1:

Input: rowIndex = 3
Output: [1,3,3,1]
Example 2:

Input: rowIndex = 0
Output: [1]
Example 3:

Input: rowIndex = 1
Output: [1,1]

Constraints:

0 <= rowIndex <= 33

Follow up: Could you optimize your algorithm to use only O(rowIndex) extra space?
*/
package main

import "fmt"

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}

	if rowIndex == 1 {
		return []int{1, 1}
	}
	last := []int{1, 1}

	for i := 2; i <= rowIndex; i++ {
		var row = []int{}
		for j, v := range last {
			if j == 0 {
				row = append(row, v)
				continue
			}
			row = append(row, v+last[j-1])
			if j == len(last)-1 {
				row = append(row, v)
			}
		}

		if i == rowIndex {
			return row
		}
		last = row
	}

	return []int{}

}

func main() {
	fmt.Println(getRow(3))
	fmt.Println(getRow(0))
	fmt.Println(getRow(1))
}

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	n := len(nums)

	res := ArrayToBST(nums, 0, n-1)

	return res
}

func ArrayToBST(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}

	mid := (start + end) / 2
	node := &TreeNode{Val: nums[mid]}
	node.Left = ArrayToBST(nums, start, mid-1)
	node.Right = ArrayToBST(nums, mid+1, end)

	return node
}

func main() {
	fmt.Println(sortedArrayToBST([]int{-10, -3, 0, 5, 9}))
	fmt.Println(sortedArrayToBST([]int{1, 3}))
}

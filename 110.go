// 110. Balanced Binary Tree
// Attempted
// Easy
// Topics
// Companies
// Given a binary tree, determine if it is
// height-balanced
// .
//
// Example 1:
//
// Input: root = [3,9,20,null,null,15,7]
// Output: true
// Example 2:
//
// Input: root = [1,2,2,3,3,null,null,4,4]
// Output: false
// Example 3:
//
// Input: root = []
// Output: true
//
// Constraints:
//
// The number of nodes in the tree is in the range [0, 5000].
// -104 <= Node.val <= 104
package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left == nil && root.Right == nil {
		return true
	}

	if math.Abs(height(root.Left)-height(root.Right)) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func height(right *TreeNode) float64 {
	if right == nil {
		return 0
	}

	return math.Max(height(right.Left), height(right.Right)) + 1
}

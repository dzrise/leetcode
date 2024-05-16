// 145. Binary Tree Postorder Traversal
// Easy
// Topics
// Companies
// Given the root of a binary tree, return the postorder traversal of its nodes' values.
//
// Example 1:
//
// Input: root = [1,null,2,3]
// Output: [3,2,1]
// Example 2:
//
// Input: root = []
// Output: []
// Example 3:
//
// Input: root = [1]
// Output: [1]
//
// Constraints:
//
// The number of the nodes in the tree is in the range [0, 100].
// -100 <= Node.val <= 100
//
// Follow up: Recursive solution is trivial, could you do it iteratively?
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postOrder(node *TreeNode, data []int) []int {
	if node == nil {
		return data
	}

	data = postOrder(node.Left, data)
	data = postOrder(node.Right, data)
	data = append(data, node.Val)

	return data

}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	data := make([]int, 0)
	data = postOrder(root, data)
	return data
}

func main() {
	fmt.Println(postorderTraversal(&TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: nil, Right: nil}, Right: nil}}))
	fmt.Println(postorderTraversal(&TreeNode{Val: 1, Left: nil, Right: nil}))
}

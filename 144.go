// Given the root of a binary tree, return the preorder traversal of its nodes' values.
//
// Example 1:
//
// Input: root = [1,null,2,3]
// Output: [1,2,3]
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
// The number of nodes in the tree is in the range [0, 100].
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

func preOrder(node *TreeNode, data []int) []int {
	if node == nil {
		return data
	}

	data = append(data, node.Val)
	data = preOrder(node.Left, data)
	data = preOrder(node.Right, data)

	return data

}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	data := make([]int, 0)
	data = preOrder(root, data)
	return data
}

func main() {
	fmt.Println(preorderTraversal(&TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: nil, Right: nil}, Right: nil}}))
	fmt.Println(preorderTraversal(&TreeNode{Val: 1, Left: nil, Right: nil}))
}

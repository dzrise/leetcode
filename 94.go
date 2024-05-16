package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inOrder(node *TreeNode, data []int) []int {
	if node == nil {
		return data
	}

	data = inOrder(node.Left, data)
	data = append(data, node.Val)
	data = inOrder(node.Right, data)
	return data

}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	data := make([]int, 0)

	data = inOrder(root, data)
	return data
}

func main() {
	fmt.Println(inorderTraversal(&TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: nil, Right: nil}, Right: nil}}))
	fmt.Println(inorderTraversal(&TreeNode{Val: 1, Left: nil, Right: nil}))
	fmt.Println(inorderTraversal(&TreeNode{}))
}

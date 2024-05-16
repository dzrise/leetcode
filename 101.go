// 101. Symmetric Tree
// Easy
// Topics
// Companies
// Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).
//
// Example 1:
//
// Input: root = [1,2,2,3,4,4,3]
// Output: true
// Example 2:
//
// Input: root = [1,2,2,null,3,null,3]
// Output: false
//
// Constraints:
//
// The number of nodes in the tree is in the range [1, 1000].
// -100 <= Node.val <= 100
//
// Follow up: Could you solve it both recursively and iteratively?
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//итеральный способ

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}

	var q []*TreeNode
	q = append(q, root.Left)
	q = append(q, root.Right)

	var node1, node2 *TreeNode

	for len(q) > 0 {
		node1 = q[len(q)-1]
		node2 = q[len(q)-2]
		q = q[:len(q)-2]

		if node1 == nil && node2 == nil {
			continue
		}

		if node1 == nil || node2 == nil {
			return false
		}

		if node1.Val != node2.Val {
			return false
		}

		q = append(q, node1.Left)
		q = append(q, node2.Right)
		q = append(q, node1.Right)
		q = append(q, node2.Left)

	}

	return true

}

// рекурсия
//func isMirror(p *TreeNode, q *TreeNode) bool {
//	if p == nil && q == nil {
//		return true
//	}
//
//	if p == nil || q == nil {
//		return false
//	}
//
//	return p.Val == q.Val && isMirror(p.Left, q.Right) && isMirror(p.Right, q.Left)
//
//}
//
//func isSymmetric(root *TreeNode) bool {
//	if root == nil {
//		return false
//	}
//
//	if root.Left == nil && root.Right == nil {
//		return true
//	}
//
//	if root.Left == nil || root.Right == nil {
//		return false
//	}
//
//	if root.Left.Val != root.Right.Val {
//		return false
//	}
//
//	return isMirror(root.Left, root.Right)
//
//}

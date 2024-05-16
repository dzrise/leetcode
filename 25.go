// 25. Reverse Nodes in k-Group
// Hard
// Topics
// Companies
// Given the head of a linked list, reverse the nodes of the list k at a time, and return the modified list.
//
// k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes, in the end, should remain as it is.
//
// You may not alter the values in the list's nodes, only nodes themselves may be changed.
//
// Example 1:
//
// Input: head = [1,2,3,4,5], k = 2
// Output: [2,1,4,3,5]
// Example 2:
//
// Input: head = [1,2,3,4,5], k = 3
// Output: [3,2,1,4,5]
//
// Constraints:
//
// The number of nodes in the list is n.
// 1 <= k <= n <= 5000
// 0 <= Node.val <= 1000
//
// Follow-up: Can you solve the problem in O(1) extra memory space?
package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	dummy := &ListNode{Val: 0, Next: head}
	prev := dummy
	count := 0
	let := head

	for let != nil {
		count++
		let = let.Next
	}

	for i := 0; i < count/k; i++ {
		for j := 0; j < k-1; j++ {
			temp := prev.Next
			prev.Next = head.Next
			head.Next = head.Next.Next
			prev.Next.Next = temp
		}
		prev = head
		head = head.Next
	}

	return dummy.Next

}

//func reverseKGroup(head *ListNode, k int) *ListNode {
//	if head == nil {
//		return nil
//	}
//
//	if head.Next == nil {
//		return head
//	}
//
//	current := head
//	var prev, next *ListNode
//	count := 0
//
//	for current != nil && count < k {
//		next = current.Next
//		current.Next = prev
//		prev = current
//		current = next
//		count++
//	}
//
//	if next != nil {
//		head.Next = reverseKGroup(next, k)
//	}
//	return prev
//}

func main() {
	fmt.Println(reverseKGroup(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}, 2))
	fmt.Println(reverseKGroup(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}, 3))
}

// 24. Swap Nodes in Pairs
// Medium
// Topics
// Companies
// Given a linked list, swap every two adjacent nodes and return its head. You must solve the problem without modifying the values in the list's nodes (i.e., only nodes themselves may be changed.)
//
// Example 1:
//
// Input: head = [1,2,3,4]
// Output: [2,1,4,3]
// Example 2:
//
// Input: head = []
// Output: []
// Example 3:
//
// Input: head = [1]
// Output: [1]
//
// Constraints:
//
// The number of nodes in the list is in the range [0, 100].
// 0 <= Node.val <= 100
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

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{}
	first := dummy
	second := dummy
	prev := dummy
	prev.Next = head
	for {
		if prev.Next == nil || prev.Next.Next == nil {
			break
		}

		first = prev.Next
		second = first.Next
		prev.Next = second
		first.Next = second.Next
		second.Next = first
		prev = second.Next
	}

	return dummy.Next
}

func main() {
	fmt.Println(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}})
}

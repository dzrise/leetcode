// 19. Remove Nth Node From End of List
// Medium
// Topics
// Companies
// Hint
// Given the head of a linked list, remove the nth node from the end of the list and return its head.
//
// Example 1:
//
// Input: head = [1,2,3,4,5], n = 2
// Output: [1,2,3,5]
// Example 2:
//
// Input: head = [1], n = 1
// Output: []
// Example 3:
//
// Input: head = [1,2], n = 1
// Output: [1]
//
// Constraints:
//
// The number of nodes in the list is sz.
// 1 <= sz <= 30
// 0 <= Node.val <= 100
// 1 <= n <= sz
//
// Follow up: Could you do this in one pass?
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func GetLenghtList(head *ListNode) int {
	tmp := head
	count := 0

	for tmp.Next != nil {
		count++
		tmp = tmp.Next
	}

	return count
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n == 0 {
		return head
	}
	if head.Next == nil {
		return nil
	}

	cnt := GetLenghtList(head)
	cntBegin := cnt - n + 2
	prev := &ListNode{}
	tmp := head

	for i := 1; i < cntBegin; i++ {
		prev = tmp
		tmp = tmp.Next
	}

	if prev.Next == nil {
		head = head.Next
		return head
	} else {
		if prev.Next != nil {
			prev.Next = prev.Next.Next
		}

		return head
	}

}

func main() {
	//fmt.Println(removeNthFromEnd(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}, 2))
	fmt.Println(removeNthFromEnd(&ListNode{Val: 1, Next: &ListNode{Val: 2}}, 2))
	//fmt.Println(removeNthFromEnd(&ListNode{Val: 1, Next: &ListNode{Val: 2}}, 1))
}

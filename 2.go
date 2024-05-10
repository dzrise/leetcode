//You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.
//
//You may assume the two numbers do not contain any leading zero, except the number 0 itself.
//
//
//
//Example 1:
//
//
//Input: l1 = [2,4,3], l2 = [5,6,4]
//Output: [7,0,8]
//Explanation: 342 + 465 = 807.
//Example 2:
//
//Input: l1 = [0], l2 = [0]
//Output: [0]
//Example 3:
//
//Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
//Output: [8,9,9,9,0,0,0,1]
//
//
//Constraints:
//
//The number of nodes in each linked list is in the range [1, 100].
//0 <= Node.val <= 9
//It is guaranteed that the list represents a number that does not have leading zeros.

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addToListNode(node *ListNode, val int) *ListNode {
	if node == nil {
		return &ListNode{Val: val}
	}
	node.Next = addToListNode(node.Next, val)

	return node
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	head := new(ListNode)
	cur := head
	for l1 != nil || l2 != nil || carry != 0 {
		n1, n2 := 0, 0
		if l1 != nil {
			n1, l1 = l1.Val, l1.Next
		}
		if l2 != nil {
			n2, l2 = l2.Val, l2.Next
		}
		num := n1 + n2 + carry
		carry = num / 10
		cur.Next = &ListNode{num % 10, nil}
		cur = cur.Next
	}
	return head.Next
}

func main() {
	list1 := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	list2 := []int{5, 6, 4}
	var l1 *ListNode
	var l2 *ListNode
	for _, v := range list1 {
		l1 = addToListNode(l1, v)
	}
	for _, v := range list2 {
		l2 = addToListNode(l2, v)
	}
	r := addTwoNumbers(l1, l2)

	current := r
	for current != nil {
		fmt.Println(current.Val) //print the added values
		current = current.Next
	}
}

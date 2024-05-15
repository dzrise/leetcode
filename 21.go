// You are given the heads of two sorted linked lists list1 and list2.
//
// Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.
//
// Return the head of the merged linked list.
//
// Example 1:
//
// Input: list1 = [1,2,4], list2 = [1,3,4]
// Output: [1,1,2,3,4,4]
// Example 2:
//
// Input: list1 = [], list2 = []
// Output: []
// Example 3:
//
// Input: list1 = [], list2 = [0]
// Output: [0]
//
// Constraints:
//
// The number of nodes in both lists is in the range [0, 50].
// -100 <= Node.val <= 100
// Both list1 and list2 are sorted in non-decreasing order.
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func GetLenghtList(head *ListNode) int {
	tmp := head
	count := 1

	for tmp.Next != nil {
		count++
		tmp = tmp.Next
	}

	return count
}

func addToListNode(node *ListNode, val int) *ListNode {
	if node == nil {
		return &ListNode{Val: val}
	}
	node.Next = addToListNode(node.Next, val)

	return node
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	n := GetLenghtList(list1) + GetLenghtList(list2)

	var res *ListNode
	for i := 0; i < n; i++ {
		if list1 == nil && list2 == nil {
			break
		}
		n1 := -101
		if list1 != nil {
			n1 = list1.Val
		}
		n2 := -101
		if list2 != nil {
			n2 = list2.Val
		}

		if n1 != -101 && n1 < n2 {
			res = addToListNode(res, list1.Val)
			list1 = list1.Next
		} else if n2 != -101 && n1 > n2 {
			res = addToListNode(res, list2.Val)
			list2 = list2.Next
		} else {
			if n1 != -101 {
				res = addToListNode(res, list1.Val)
				list1 = list1.Next
			}
			if n2 != -101 {
				res = addToListNode(res, list2.Val)
				list2 = list2.Next
			}
		}
	}
	return res
}

func main() {
	//fmt.Println(mergeTwoLists(nil, nil))
	fmt.Println(mergeTwoLists(&ListNode{Val: 1}, &ListNode{Val: 2}))
	//fmt.Println(mergeTwoLists(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}, &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}))
}

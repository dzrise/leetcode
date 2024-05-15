// You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.
//
// Merge all the linked-lists into one sorted linked-list and return it.
//
// Example 1:
//
// Input: lists = [[1,4,5],[1,3,4],[2,6]]
// Output: [1,1,2,3,4,4,5,6]
// Explanation: The linked-lists are:
// [
// 1->4->5,
// 1->3->4,
// 2->6
// ]
// merging them into one sorted list:
// 1->1->2->3->4->4->5->6
// Example 2:
//
// Input: lists = []
// Output: []
// Example 3:
//
// Input: lists = [[]]
// Output: []
package main

import (
	"fmt"
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	vals := make([]int, 0, 10000)

	for _, v := range lists {
		for v != nil {
			vals = append(vals, v.Val)
			v = v.Next
		}
	}

	if len(vals) == 0 {
		return nil
	}

	sort.Ints(vals)
	var res, curr *ListNode

	for _, val := range vals {
		node := &ListNode{val, nil}
		if res == nil {
			res = node
			curr = node
			continue
		}
		curr.Next = node
		curr = node
	}

	return res
}

func main() {
	fmt.Println(mergeKLists([]*ListNode{&ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}, &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}, &ListNode{Val: 2, Next: &ListNode{Val: 6}}}))
}

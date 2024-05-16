package main

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

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}
	sl := map[int]bool{head.Val: true}
	res := &ListNode{Val: head.Val}

	for head != nil {
		if !sl[head.Val] {
			addToListNode(res, head.Val)
		}
		sl[head.Val] = true
		head = head.Next
	}
	return res
}

package main

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

func deleteDuplicates(head *ListNode) *ListNode {
	tail := head

	for tail != nil && tail.Next != nil {
		if tail.Val == tail.Next.Val {
			tail.Next = tail.Next.Next
		} else {
			tail = tail.Next
		}
	}

	return head
}

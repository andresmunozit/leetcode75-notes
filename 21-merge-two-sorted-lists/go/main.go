package main

import (
	helpersgo "github.com/andresmunozit/leetcode75-notes/helpers-go"
)

func mergeTwoLists(list1 *helpersgo.ListNode, list2 *helpersgo.ListNode) *helpersgo.ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	dummy := &helpersgo.ListNode{
		Val:  0,
		Next: nil,
	}

	tail := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}

	if list1 != nil {
		tail.Next = list1
	}

	if list2 != nil {
		tail.Next = list2
	}

	return dummy.Next
}

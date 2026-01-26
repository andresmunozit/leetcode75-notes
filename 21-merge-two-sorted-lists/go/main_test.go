package main

import (
	"testing"

	helpersgo "github.com/andresmunozit/leetcode75-notes/helpers-go"
	"github.com/stretchr/testify/require"
)

func Test_mergeTwoLists(t *testing.T) {
	tests := []struct {
		name     string
		list1    *helpersgo.ListNode
		list2    *helpersgo.ListNode
		expected *helpersgo.ListNode
	}{
		{
			name:     "two lists of equal length",
			list1:    helpersgo.SliceToLinkedList([]int{1, 2, 4}),
			list2:    helpersgo.SliceToLinkedList([]int{1, 3, 4}),
			expected: helpersgo.SliceToLinkedList([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			name:     "both lists are empty",
			list1:    helpersgo.SliceToLinkedList([]int{}),
			list2:    helpersgo.SliceToLinkedList([]int{}),
			expected: nil,
		},
		{
			name:     "list1 is empty",
			list1:    helpersgo.SliceToLinkedList([]int{}),
			list2:    helpersgo.SliceToLinkedList([]int{0}),
			expected: helpersgo.SliceToLinkedList([]int{0}),
		},
		{
			name:     "list2 is empty",
			list1:    helpersgo.SliceToLinkedList([]int{0}),
			list2:    helpersgo.SliceToLinkedList([]int{}),
			expected: helpersgo.SliceToLinkedList([]int{0}),
		},
		{
			name:     "list1 is shorter",
			list1:    helpersgo.SliceToLinkedList([]int{20, 25}),
			list2:    helpersgo.SliceToLinkedList([]int{15, 41, 45, 52}),
			expected: helpersgo.SliceToLinkedList([]int{15, 20, 25, 41, 45, 52}),
		},
		{
			name:     "list2 is shorter",
			list1:    helpersgo.SliceToLinkedList([]int{12, 53, 54, 88}),
			list2:    helpersgo.SliceToLinkedList([]int{0, 12}),
			expected: helpersgo.SliceToLinkedList([]int{0, 12, 12, 53, 54, 88}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, mergeTwoLists(tt.list1, tt.list2))
		})
	}
}

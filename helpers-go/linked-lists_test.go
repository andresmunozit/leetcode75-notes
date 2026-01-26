package helpersgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceToLinkedList(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		expected *ListNode
	}{
		{
			name:  "success case",
			slice: []int{1, 4, 6, 5, 7},
			expected: func() *ListNode {
				head := &ListNode{
					Val:  1,
					Next: nil,
				}

				tail := head
				tail.Next = &ListNode{
					Val:  4,
					Next: nil,
				}
				tail = tail.Next

				tail.Next = &ListNode{
					Val:  6,
					Next: nil,
				}
				tail = tail.Next

				tail.Next = &ListNode{
					Val:  5,
					Next: nil,
				}
				tail = tail.Next

				tail.Next = &ListNode{
					Val:  7,
					Next: nil,
				}

				return head
			}(),
		},
		{
			name:     "empty slice case",
			slice:    []int{},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, SliceToLinkedList(tt.slice))
		})
	}
}

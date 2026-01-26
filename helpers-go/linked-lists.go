package helpersgod

type ListNode struct {
	Val  int
	Next *ListNode
}

func SliceToLinkedList(slice []int) *ListNode {
	dummy := &ListNode{}

	tail := dummy

	for i := 0; i < len(slice); i++ {
		tail.Next = &ListNode{
			Val:  slice[i],
			Next: nil,
		}
		tail = tail.Next
	}

	return dummy.Next
}

func LinkedListToSlice(ll *ListNode) []int {
	slice := []int{}

	for ll != nil {
		slice = append(slice, ll.Val)
		ll = ll.Next
	}

	return slice
}

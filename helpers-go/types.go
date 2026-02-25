package helpersgo

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int       `json:"val"`
	Left  *TreeNode `json:"left,omitempty"`
	Right *TreeNode `json:"right,omitempty"`
}

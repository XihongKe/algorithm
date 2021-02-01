package golang

type ListNode struct {
	Val  int
	Next *ListNode
}

func ListNodeConstructor(nums []int) *ListNode {
	listNode := &ListNode{nums[0], nil}
	tmp := listNode
	for i := 1; i < len(nums); i++ {
		tmp.Next = &ListNode{nums[i], nil}
		tmp = tmp.Next
	}
	return listNode
}

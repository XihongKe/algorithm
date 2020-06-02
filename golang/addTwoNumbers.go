package golang

/**
问题：给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

由于数字的顺序是倒过来的，只需要做普通的加法运算（同位相加，大于9进1）即可
 */
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 假头
	dummyHead := &ListNode{0, nil}
	// 链表的指针
	cur := dummyHead
	// 进位的值，比如当前指针l1.val = 8,l2.val = 9, sum = l1.val+l2.val=17，需要进一位，则carry=1
	carry := 0
	for l1 != nil || l2 != nil {
		// 计算当前位
		sum := carry
		if l1 != nil {
			sum += l1.Val
		}
		if l2 != nil {
			sum += l2.Val
		}
		// 判断是否需要进位
		if sum >= 10 {
			sum -= 10
			carry = 1
		} else {
			carry = 0
		}
		cur.Next = &ListNode{sum, nil}

		// 指针迭代
		cur = cur.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	// 有两个链表都走完了，但是还有进位的情况，比如l1 = {5, nil}, l2 = {6, nil}
	if carry != 0 {
		cur.Next = &ListNode{carry, nil}
	}
	return dummyHead.Next
}
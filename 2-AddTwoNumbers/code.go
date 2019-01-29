package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		root  = &ListNode{}
		curr  = root
		carry = 0
	)

	for l1 != nil || l2 != nil || carry != 0 {
		val := carry

		if l1 != nil {
			val += l1.Val
		}

		if l2 != nil {
			val += l2.Val
		}

		carry, val = val/10, val%10

		curr.Next = &ListNode{Val: val}
		curr = curr.Next

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}

	return root.Next
}

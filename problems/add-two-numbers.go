package problems

// https://leetcode.com/problems/add-two-numbers


type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := &ListNode{}
	head := sum

	add := 0
	n := 0
	v1, v2 := 0, 0

	for l1 != nil || l2 != nil {
		if l1 == nil {
			v1 = 0
		} else {
			v1 = l1.Val
		}
		if l2 == nil {
			v2 = 0
		} else {
			v2 = l2.Val
		}
		//

		n = v1 + v2 + add
		add = int(n/10)
		if add > 0 {
			sum.Val = n % 10
		} else {
			sum.Val = n
		}

		if (l2 != nil && l2.Next != nil) || (l1 != nil && l1.Next != nil) {
			next := &ListNode{}
			sum.Next = next
			sum = next
		}

		//
		if l2 != nil {
			l2 = l2.Next
		}
		if l1 != nil {
			l1 = l1.Next
		}
		n = 0
	}

	if add > 0 {
		sum.Next = &ListNode{
			Val: add,
		}
	}

	return head
}

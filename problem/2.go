package problem

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbersWithOne(l1, l2, 0)

}

func addTwoNumbersWithOne(l1 *ListNode, l2 *ListNode, v int) *ListNode {
	h := &ListNode{}
	if l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + v
		h.Val = sum % 10
		h.Next = addTwoNumbersWithOne(l1.Next, l2.Next, sum / 10)
	} else if l1 == nil && l2 == nil  {
		if v > 0 {
			h.Val = v
		} else {
			return nil
		}
	} else if l1 == nil && l2 != nil {
		h.Val = l2.Val + v
		h.Next = l2.Next
	} else if l2 == nil && l1 != nil {
		h.Val = l1.Val + v
		h.Next = l1.Next
	}
	return h
}

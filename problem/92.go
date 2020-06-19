package problem

func ReverseBetween(head *ListNode, m int, n int) *ListNode {
	i := 1
	var pre *ListNode
	curr := head
	for i < m {
		pre = curr
		curr = curr.Next
		i++
	}
	curr2 := curr
	next := curr.Next
	for next != nil && i < n {
		next2 := next.Next
		next.Next = curr
		curr = next
		next = next2
		i ++
	}
	if next != nil {
		curr2.Next = next
	} else {
		curr2.Next = nil
	}

	if m == 1 {
		return curr
	} else {
		pre.Next = curr
		return head
	}

}
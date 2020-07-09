package problem

func ReorderList(head *ListNode) {
	if head == nil {
		return
	}
	fast := head
	slow := head
	dummy := head
	for fast != nil {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
			if fast == nil {
				break
			}
		} else {
			break
		}
		slow = slow.Next
	}
	curr := slow.Next
	if curr == nil {
		return
	}
	slow.Next = nil
	next := curr.Next
	curr.Next = nil
	for next != nil {
		n := next.Next
		next.Next = curr
		curr = next
		next = n
	}
	for curr != nil {
		n := dummy.Next
		m := curr.Next
		dummy.Next = curr
		curr.Next = n
		dummy = n
		curr = m
	}
}

package problem

func swapPairs(head *ListNode) *ListNode {
	if (head == nil) || (head.Next == nil) {
		return head
	}
	h := ListNode{Val: 0, Next: head}
	prev := &h
	curr := head
	next := head.Next
	for {
		t := next.Next
		next.Next = curr
		prev.Next = next
		curr.Next = t
		prev = curr
		curr = curr.Next
		if curr == nil{
			break
		}
		next = curr.Next
		if next == nil {
			break
		}
	}
	return h.Next
}
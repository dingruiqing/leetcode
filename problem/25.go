package problem

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	h := &ListNode{Val: 0, Next: head}
	prev := h
	curr := head
	next := curr.Next
	t := head
	for c := 1; c <=k; c++ {
		t = t.Next
		if t == nil {
			return head
		}
	}
	for i := 1; i < k; i++ {
		n := next
		next = n.Next
		n.Next = curr
		curr = n
	}
	prev.Next = curr
	head.Next = reverseKGroup(next, k)
	return h.Next
}
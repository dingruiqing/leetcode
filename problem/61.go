package problem

func RotateRight(head *ListNode, k int) *ListNode {
	if head == nil {return head}
	fast := head
	slow := head
	for k > 0 {
		fast = fast.Next
		if fast == nil {
			fast = head
		}
		k--
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	a := slow.Next
	if a == nil {
		return head
	}
	slow.Next = nil
	fast.Next = head
	return a
}

package problem

func IsPalindrome(head *ListNode) bool {
	size := 0
	curr := head
	for curr != nil {
		curr = curr.Next
		size ++
	}
	if size == 0 || size == 1 {
		return true
	}
	a := size / 2 - 1
	b := a + 1
	curr = head
	next := curr.Next
	for a > 0 {
		t := next.Next
		next.Next = curr
		curr = next
		next = t
		a --
	}
	if size % 2 != 0 {
		next = next.Next
	}
	for b > 0 {
		if curr.Val != next.Val {
			return false
		}
		curr = curr.Next
		next = next.Next
		b --
	}
	return true
}

package problem

import "math"

func InsertionSortList(head *ListNode) *ListNode {

	dummy := &ListNode{math.MinInt32, head}
	curr := head
	end := head
	c := dummy
	for curr!= nil {
		next := curr.Next
		if curr.Val <= c.Val {c = dummy}
		for c.Next != nil && curr.Val >= c.Next.Val {
			if c == end {
				break
			}
			c = c.Next
		}
		if c == end {
			end = curr
			curr = curr.Next
			continue
		}
		end.Next = curr.Next
		curr.Next = c.Next
		c.Next = curr
		curr = next
	}
	return dummy.Next
}

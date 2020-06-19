package problem

import "math"

func Partition(head *ListNode, x int) *ListNode {
	a := &ListNode{math.MinInt32, head}
	pre := a
	insert := a
	curr := head
	for curr != nil {
		if curr.Val < x {
			next := curr.Next
			pre.Next = next
			curr.Next = insert.Next
			insert.Next = curr
			insert = curr
			if pre.Next == insert {
				pre = pre.Next
			}
			curr = next
		} else {
			pre = curr
			curr = curr.Next
		}
	}
	return a.Next
}

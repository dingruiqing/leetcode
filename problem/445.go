package problem

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	a, b := []int{}, []int{}
	for l1 != nil {
		a = append(a, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		b = append(b, l2.Val)
		l2 = l2.Next
	}
	i := len(a) - 1
	j := len(b) - 1
	var d *ListNode
	next := 0
	for {
		c := &ListNode{}
		sum := 0
		if i >= 0 && j >= 0 {
			sum = a[i] + b[j] + next
		} else if i >= 0 && j < 0 {
			sum = a[i] + next
		} else if j >= 0 && i < 0 {
			sum = b[j] + next
		} else {
			if next == 0 {
				break
			}
			sum = next
		}
		c.Val = sum % 10
		next = sum / 10
		c.Next = d
		d = c
		i --
		j --
	}
	return d
}

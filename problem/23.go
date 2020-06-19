package problem

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	} else if len(lists) == 2 {
		node1 := lists[0]
		node2 := lists[1]
		if (node1 == nil) && (node2 == nil) {
			return nil
		} else if node1 == nil {
			return node2
		} else if node2 == nil {
			return node1
		}
		nodeHead := ListNode{Val: 0}
		n := &nodeHead
		for {
			if node1.Val <= node2.Val {
				n.Next = node1
				n = node1
				if n.Next != nil {
					node1 = node1.Next
				} else {
					n.Next = node2
					return nodeHead.Next
				}
			} else {
				n.Next = node2
				n = node2
				if n.Next != nil {
					node2 = node2.Next
				} else {
					n.Next = node1
					return nodeHead.Next
				}
			}
		}
	} else {
		mid := len(lists) / 2
		a := mergeKLists(lists[:mid])
		b := mergeKLists(lists[mid:])
		var c = []*ListNode{a,b}
		return mergeKLists(c)
	}

}

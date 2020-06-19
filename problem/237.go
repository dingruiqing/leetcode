package problem

func DeleteNode(node *ListNode) {
	for {
		if node.Next == nil {
			node = nil
			break
		}
		node.Val = node.Next.Val
		node = node.Next
	}
}

package problem

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func BuildListNode(nums []int) *ListNode {
	var b *ListNode
	for i := len(nums) - 1; i >= 0; i-- {
		a := &ListNode{nums[i], b}
		b = a
	}
	return b
}

func PrintListNode(node *ListNode) {
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

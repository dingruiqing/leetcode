package problem

type Node struct {
	Val string
	Index int
	Next *Node
}


func longestValidParentheses(s string) int {
	head := &Node{Index: -1}
	stack := &Node{Index: len(s), Next: head}
	for i, v := range s {
		if stack.Next.Val == "(" && string(v) == ")" {
			stack.Next = stack.Next.Next
		} else {
			t := Node{Val: string(v), Index: i}
			t.Next = stack.Next
			stack.Next = &t
		}
	}
	c := 0
	for {
		v := stack.Index
		t := stack.Next
		if t != nil {
			if v - t.Index - 1 > c {
				c = v - t.Index - 1
			}
			stack = t
		} else {
			return c
		}
	}
}
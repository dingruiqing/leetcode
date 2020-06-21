package problem

type Stack struct {
	top    *ListNode
	length int
}

func New() *Stack {
	return &Stack{nil, 0}
}

func (s *Stack) Peek() (int, bool) {
	if s.length == 0 {
		return 0, false
	}
	return s.top.Val, true
}

func (s *Stack) Push(i int) {
	a := &ListNode{i, s.top}
	s.top = a
	s.length++
}

func (s *Stack) Pop() (int, bool) {
	if s.length == 0 {
		return 0, false
	}
	a := s.top
	s.top = a.Next
	s.length--
	return a.Val, true
}

func (s *Stack) Len() int {
	return s.length
}

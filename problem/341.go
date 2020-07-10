package problem

//type NestedInteger struct {}
//func (this NestedInteger) IsInteger() bool {}
//func (this NestedInteger) GetInteger() int {}
//func (n *NestedInteger) SetInteger(value int) {}
//func (this *NestedInteger) Add(elem NestedInteger) {}
//func (this NestedInteger) GetList() []*NestedInteger {}
//
//
//type NestedIterator struct {
//	s *Stack
//
//}
//
//func Constructor(nestedList []*NestedInteger) *NestedIterator {
//	l := len(nestedList)
//	s := NewStack()
//	for i := l - 1; i >= 0; i--{
//		PushInteger(nestedList[i], s)
//	}
//	n := &NestedIterator{s}
//	return n
//}
//
//func PushInteger(integer *NestedInteger, stack *Stack) {
//	if !integer.IsInteger() {
//		l := len(integer.GetList())
//		for i := l - 1; i >= 0; i--{
//			PushInteger(integer.GetList()[i], stack)
//		}
//	} else {
//		stack.Push(integer)
//	}
//}
//
//func (this *NestedIterator) Next() int {
//	a := this.s.Pop()
//	return a.GetInteger()
//}
//
//func (this *NestedIterator) HasNext() bool {
//	if this.s.Len() == 0 {
//		return false
//	} else {
//		return true
//	}
//}
//
//
//type Stack struct {
//	top *Node
//	length int
//}
//
//type Node struct {
//	Val  *NestedInteger
//	Next *Node
//}
//
//func NewStack() *Stack {
//	return &Stack{nil, 0}
//}
//
//func (s *Stack) Peek() *NestedInteger {
//	return s.top.Val
//}
//
//func (s *Stack) Push(i *NestedInteger) {
//	a := &Node{i, s.top}
//	s.top = a
//	s.length ++
//}
//
//func (s *Stack) Pop() *NestedInteger {
//	a := s.top
//	s.top = a.Next
//	s.length --
//	return a.Val
//}
//
//func (s *Stack) Len() int {
//	return s.length
//}

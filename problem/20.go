package problem

import "github.com/golang-collections/collections/stack"

func IsValid(s string) bool {
	st := stack.New()
	for _, v := range s {
		str := string(v)
		if str == "(" || str == "{" || str == "[" {
			st.Push(v)
		} else if str == ")" || str == "}" || str == "]" {
			top := st.Peek()
			if (str == ")" && top == v - 1) || (str != ")" && top == v - 2) {
				st.Pop()
			} else {
				return false
			}
		}
	}
	if st.Len() != 0 {
		return false
	}
	return true
}

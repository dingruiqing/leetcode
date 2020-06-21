package problem

func IsValid(s string) bool {
	st := New()
	for _, v := range s {
		str := string(v)
		if str == "(" || str == "{" || str == "[" {
			st.Push(int(v))
		} else if str == ")" || str == "}" || str == "]" {
			top, ok := st.Peek()
			if ok {
				if (str == ")" && top == int(v-1)) || (str != ")" && top == int(v-2)) {
					st.Pop()
				} else {
					return false
				}
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

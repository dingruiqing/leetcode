package problem

//func SimplifyPath(path string) string {
//	stack := NewStack()
//	paths:= strings.Split(path, "/")
//	for _, v := range paths {
//		if v == ".." {
//			stack.Pop()
//		} else if v != "." && v != "" {
//			stack.Push(v)
//		}
//	}
//	res := ""
//	for stack.length!= 0 {
//		p, ok := stack.Pop()
//		if ok {
//			res = "/" + p + res
//		} else {
//			break
//		}
//	}
//	return res
//}

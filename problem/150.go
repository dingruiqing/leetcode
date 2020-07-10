package problem

import "strconv"

func evalRPN(tokens []string) int {
	stack := NewStack()
	for _, v := range tokens {
		if v == "+" || v == "-" || v == "*" || v == "/" {
			a, _ := stack.Pop()
			b, _ := stack.Pop()
			switch v {
			case "+":
				stack.Push(b + a)
			case "-":
				stack.Push(b - a)
			case "*":
				stack.Push(b * a)
			case "/":
				stack.Push(b / a)
			}
		} else {
			a, _ := strconv.Atoi(v)
			stack.Push(a)
		}
	}
	res, _ := stack.Pop()
	return res
}

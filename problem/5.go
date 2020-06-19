package problem

func longestPalindrome(s string) string {
	if len(s) == 1{
		return s
	}
	if len(s) == 0 {
		return s
	}
	max := 0
	var str string
	for index, _ := range s {
		if index >= 1 && index <len(s) - 1 {
			a := checkEqual(index, index + 1, s)
			if len(a) > max{
				str = a
				max = len(a)
			}
			b := checkEqual(index - 1, index + 1, s)
			if len(b) > max {
				str = b
				max = len(b)
			}
		}
	}
	return str
}

func checkEqual(i int, j int, s string) string {
	for s[i] == s[j] {
		i -= 1
		j += 1
		if i == -1 || j == len(s){
			break
		}
	}
	return s[i + 1: j]
}

package problem

func multiply(num1 string, num2 string) string {
	l := make([]int32, len(num1) + len(num2) + 1)
	for k1, v1 := range num1 {
		for k2, v2 := range num2 {
			a := (v1 - 48) * (v2 - 48)
			l[len(num1) + len(num2) - k1 - k2 - 2] += a
		}
	}
	for k, v := range l {
		if v >= 100 {
			l[k + 2] += v / 100
			v = v % 100
		}
		if v >= 10 {
			l[k + 1] += v / 10
			v = v % 10
		}
		l[k] = v
	}
	str := ""
	i := len(l) - 1
	flag := false
	for ;i >= 0; i -- {
		if l[i] > 0 {
			flag = true
		}
		if flag{
			str = str + string(l[i] + 48)
		}
	}
	return str
}

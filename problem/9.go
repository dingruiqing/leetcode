package problem

func IsPalindrome2(x int) bool {
	if x < 0 {
		return false
	}
	i := 1
	n := x
	for n >= 10 {
		n = n / 10
		i *= 10
	}
	for x > 0 {
		a := x / i
		b := x % 10
		if x < i {
			a = 0
		}
		if a != b {
			return false
		}
		x = x % i / 10
		i /= 100
	}
	return true
}

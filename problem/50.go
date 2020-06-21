package problem

func MyPow(x float64, n int) float64 {
	switch {
	case n == 0:
		return 1
	case n > 0:
		a := MyPow(x, n/2)
		if n%2 == 0 {
			return a * a
		} else {
			return a * a * x
		}
	case n < 0:
		return 1 / MyPow(x, -n)
	}
	return 0
}

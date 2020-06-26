package problem

import "math"

func integerBreak(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	a := n / 3
	b := n % 3
	if b == 0 {
		return int(math.Pow(3, float64(a)))
	} else if b == 1 {
		return int(math.Pow(3, float64(a-1))) * 4
	} else {
		return int(math.Pow(3, float64(a))) * 2
	}
}

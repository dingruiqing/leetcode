package problem

import (
	"strconv"
)

func gcd(x int, y int) int {
	if y == 0{
		return x
	} else {
		return gcd(y, x % y)
	}
}

func maxPoints(points [][]int) int {
	l := len(points)
	res := 0
	for i := 0; i < l; i++ {
		m := make(map[string]int)
		y := 0
		x := 0
		for j := 0; j < l; j++ {
			if points[j][0] == points[i][0] && points[j][1] == points[i][1] {
				x ++
			} else if points[j][0] == points[i][0] {
				y ++
			} else {
				dy := points[j][1] - points[i][1]
				dx := points[j][0] - points[i][0]
				g := gcd(dy, dx)
				if g != 0 {
					dy /= g
					dx /= g
				}
				m[strconv.Itoa(dy) + "/" + strconv.Itoa(dx)]++
			}
		}
		if x > res{
			res = x
		}
		if y + x > res {
			res = y + x
		}
		for _, v := range m {
			if v + x > res {
				res = v + x
			}
		}
	}
	return res
}

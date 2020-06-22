package problem

func NumSquares(n int) int {
	a := []int{n}
	squareList := []int{}
	for j := 1; n >= j*j; j++ {
		squareList = append(squareList, j*j)
	}
	level := 0
	for {
		b := []int{}
		for _, value := range a {
			if value == 0 {
				return level
			}
			for _, v := range squareList {
				if v <= value {
					b = append(b, value-v)
				}
			}
		}
		level++
		a = b
	}
}

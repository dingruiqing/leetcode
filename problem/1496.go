package problem

func IsPathCrossing(path string) bool {
	m := make(map[int]map[int]bool)
	point := []int{0, 0}
	m[0] = make(map[int]bool)
	m[0][0] = true
	for _, v := range path {
		switch v {
		case 78:
			point[1]++
		case 69:
			point[0]++
		case 83:
			point[1]--
		case 87:
			point[0]--
		}
		a, ok := m[point[0]]
		if ok {
			_, ok := a[point[1]]
			if ok {
				return true
			} else {
				a[point[1]] = true
			}
		} else {
			m[point[0]] = make(map[int]bool)
			m[point[0]][point[1]] = true
		}
	}
	return false
}

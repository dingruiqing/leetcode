package problem

func UniquePaths(m int, n int) int {
	nums := make([][]int, m)
	for i, _ := range nums {
		nums[i] = make([]int, n)
	}
	path := [][]int{{0, -1}, {-1, 0}}
	queue := [][]int{{m - 1, n - 1}}
	nums[m-1][n-1] = 1
	pointMap := make(map[int]map[int]bool)
	for i := 0; i <= m; i++ {
		pointMap[i] = make(map[int]bool)
	}
	for len(queue) != 0 {
		tmp := [][]int{}
		for _, q := range queue {
			for _, v := range path {
				i := v[0] + q[0]
				j := v[1] + q[1]
				if i >= 0 && j >= 0 {
					nums[i][j] += nums[q[0]][q[1]]
					_, ok := (pointMap[i])[j]
					if !ok {
						pointMap[i][j] = true
						tmp = append(tmp, []int{i, j})
					}
				}

			}
		}
		queue = tmp
	}
	return nums[0][0]
}

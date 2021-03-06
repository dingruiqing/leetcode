package problem

//func MinPathSum(grid [][]int) int {
//	path := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
//	m := len(grid)
//	if m == 0 {
//		return 0
//	}
//	n := len(grid[0])
//	nums := make([][]int, m)
//	visit := make([][]bool, m)
//	for i := 0; i < m; i++ {
//		nums[i] = make([]int, n)
//		visit[i] = make([]bool, n)
//	}
//	nums[m-1][n-1] = grid[m-1][n-1]
//	visit[m-1][n-1] = true
//	queue := [][]int{{m - 1, n - 1}}
//	for len(queue) != 0 {
//		tmp := [][]int{}
//		pointMap := make(map[string]bool)
//		for _, p := range queue {
//			for _, v := range path {
//				i := p[0] + v[0]
//				j := p[1] + v[1]
//				if i < m && j < n && i >= 0 && j >= 0 && !visit[i][j] {
//					if nums[i][j] > nums[p[0]][p[1]]+grid[i][j] || nums[i][j] == 0 {
//						nums[i][j] = nums[p[0]][p[1]] + grid[i][j]
//					}
//					_, ok := pointMap[strconv.Itoa(i)+","+strconv.Itoa(j)]
//					if ok {
//						continue
//					} else {
//						tmp = append(tmp, []int{i, j})
//						pointMap[strconv.Itoa(i)+","+strconv.Itoa(j)] = true
//					}
//				}
//			}
//		}
//		for _, v := range tmp {
//			visit[v[0]][v[1]] = true
//		}
//		queue = tmp
//	}
//	return nums[0][0]
//}

func MinPathSum(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	for j := m - 1; j >= 0; j-- {
		for i := n - 1; i >= 0; i-- {
			if j+1 < m && i+1 < n {
				grid[j][i] = min(grid[j+1][i], grid[j][i+1]) + grid[j][i]
			} else if j+1 == m && i+1 != n {
				grid[j][i] = grid[j][i+1] + grid[j][i]
			} else if i+1 == n && j+1 != m {
				grid[j][i] = grid[j+1][i] + grid[j][i]
			}
		}
	}
	return grid[0][0]
}

package problem

func LongestIncreasingPath(matrix [][]int) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	direct := [][]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, m)
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}
	max := 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(i, j, m, n, matrix, dp, direct)
			if max < dp[i][j] {
				max = dp[i][j]
			}
		}
	}
	return max
}

func dfs(i, j, m, n int, matrix, dp, direct [][]int) {
	if dp[i][j] > 0 {
		return
	}
	value := matrix[i][j]
	max := 1
	for _, d := range direct {
		x, y := i, j
		x += d[0]
		y += d[1]
		if x >= 0 && x < m && y >= 0 && y < n {
			if matrix[x][y] < value {
				if dp[x][y] < 0 {
					dfs(x, y, m, n, matrix, dp, direct)
				}
				if max < dp[x][y]+1 {
					max = dp[x][y] + 1
				}
			}
		}
	}
	dp[i][j] = max
}

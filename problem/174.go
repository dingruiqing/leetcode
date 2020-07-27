package problem

func CalculateMinimumHP(dungeon [][]int) int {
	m := len(dungeon)
	if m == 0 {
		return 0
	}
	n := len(dungeon[0])
	dp := make([][]int, m+1)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j < n+1; j++ {
			dp[i][j] = 0x3f3f3f3f
		}
	}
	dp[m][n-1] = 1
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			dp[i][j] = max(1, min(dp[i+1][j], dp[i][j+1])-dungeon[i][j])
		}
	}
	return dp[0][0]
}

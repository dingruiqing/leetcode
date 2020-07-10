package problem

func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][3]int, n)
	for i, _ := range dp {
		dp[i] = [3]int{}
	}
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	dp[0][2] = 0
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][2]-prices[i], dp[i-1][0])
		dp[i][1] = dp[i-1][0] + prices[i]
		dp[i][2] = max(dp[i-1][1], dp[i-1][2])
	}
	return max(dp[n-1][1], dp[n-1][2])
}

package problem

func DivisorGame(N int) bool {
	if N == 1 {
		return false
	}
	if N == 2 {
		return true
	}
	dp := make([]bool, N+1)
	dp[1] = false
	dp[2] = true
	for i := 3; i <= N; i++ {
		dp[i] = false
		for j := 1; j <= i/2; j++ {
			if i%j == 0 && dp[i-j] == false {
				dp[i] = true
				break
			}
		}
	}
	return dp[N]
}

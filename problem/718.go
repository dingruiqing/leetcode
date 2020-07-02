package problem

func findLength(A []int, B []int) int {
	lenA, lenB := len(A), len(B)
	dp := make([][]int, lenA)
	for i, _ := range dp {
		dp[i] = make([]int, lenB)
	}
	ans := 0
	for i := 1; i < lenA; i++ {
		for j := 1; j < lenB; j++ {
			if A[i] == B[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = 0
			}
			if ans < dp[i][j] {
				ans = dp[i][j]
			}
		}
	}
	return ans
}

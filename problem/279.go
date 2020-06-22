package problem

func NumSquares(n int) int {
	nums := make([]int, n+1)
	for i, _ := range nums {
		nums[i] = n
	}
	nums[0] = 0
	for i := 1; i <= n; i++ {
		for j := 1; i >= j*j; j++ {
			if nums[i-j*j]+1 < nums[i] {
				nums[i] = nums[i-j*j] + 1
			}
		}
	}
	return nums[n]
}

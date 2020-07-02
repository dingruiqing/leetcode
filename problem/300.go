package problem

func LengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	res := make([]int, n)
	max := 0
	for i := 0; i < n; i++ {
		res[i] = 1
		for j := 0; j <= i; j++ {
			if nums[i] > nums[j] {
				v := res[j] + 1
				if v > res[i] {
					res[i] = v
				}
			}
		}
		if res[i] > max {
			max = res[i]
		}
	}
	return max
}

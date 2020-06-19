package problem

func minSubArrayLen(s int, nums []int) int {
	start := -1
	end := 0
	n := len(nums)
	minLen := n + 1
	sum := nums[0]
	for end < n {
		if sum >= s {
			if end - start < minLen {
				minLen = end - start
			}
			start ++
			sum = sum - nums[start]
		} else {
			end ++
			sum = sum + nums[end]
		}
	}
	if minLen == n + 1{
		minLen = 0
	}
	return minLen
}

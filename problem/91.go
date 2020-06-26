package problem

import "strconv"

func NumDecodings(s string) int {
	n := len(s)
	nums := make([]int, n+1)
	nums[n] = 1
	if s[n-1] != 48 {
		nums[n-1] = 1
	} else {
		nums[n-1] = 0
	}
	for i := n - 2; i >= 0; i-- {
		v1 := s[i] - 48
		v2, _ := strconv.Atoi(s[i : i+2])
		if v1 != 0 && v2 <= 26 {
			nums[i] = nums[i+1] + nums[i+2]
		} else if v1 == 0 {
			if i-1 < 0 || s[i-1] <= 48 || s[i-1] >= 51 {
				return 0
			} else {
				nums[i] = 0
			}
		} else {
			nums[i] = nums[i+1]
		}
	}
	return nums[0]
}

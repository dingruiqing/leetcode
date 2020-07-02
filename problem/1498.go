package problem

import (
	"sort"
)

func NumSubseq(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	start := 0
	end := n - 1
	sum := 0
	pow := map[int]int{0: 1}
	for i := 1; i < n; i++ {
		pow[i] = (pow[i-1] << 1) % (1e9 + 7)
	}
	for start <= end {
		if nums[start]+nums[end] > target {
			end--
		} else {
			sum += pow[end-start]
			start++
		}
	}
	return sum % (1e9 + 7)
}

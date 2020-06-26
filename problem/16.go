package problem

import "sort"

func ThreeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	if nums[0]+nums[1]+nums[2]-target > 0 {
		return nums[0] + nums[1] + nums[2]
	}
	if nums[n-1]+nums[n-2]+nums[n-3]-target < 0 {
		return nums[n-1] + nums[n-2] + nums[n-3]
	}

	min := target
	if target < 0 {
		min = -target - 1
	} else {
		min += 1
	}
	res := 0
	for i, v := range nums {
		j := i + 1
		k := n - 1
		t := target - v
		for j < k {
			d := t - nums[j] - nums[k]
			if d > 0 {
				if d < min {
					min = d
					res = target - d
				}
				j++
			} else if d < 0 {
				if -d < min {
					min = -d
					res = target - d
				}
				k--
			} else {
				return 0
			}
		}
	}
	return res
}

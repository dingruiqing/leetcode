package problem

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	res := make([][]int, 0)
	if n == 0 {
		return res
	}
	for i := 0; i <= n-3; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		j := n - 1
		k := i + 1
		t := -nums[i]
		for k < j {
			if nums[j]+nums[k] > t {
				j--
			} else if nums[j]+nums[k] < t {
				k++
			} else {
				res = append(res, []int{nums[i], nums[k], nums[j]})
				tmp := nums[j]
				for tmp == nums[j] && j > k {
					j--
				}
				tmp = nums[k]
				for tmp == nums[k] && j > k {
					k++
				}
			}
		}
	}
	return res
}

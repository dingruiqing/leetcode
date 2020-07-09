package problem

func ContainsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]bool)
	n := len(nums)
	for i := 0; i <= k && i <= n-1; i++ {
		_, ok := m[nums[i]]
		if ok {
			return true
		} else {
			m[nums[i]] = true
		}
		if i == n {
			return false
		}
	}
	i := 0
	j := k + 1
	for j < n {
		delete(m, nums[i])
		_, ok := m[nums[j]]
		if ok {
			return true
		} else {
			m[nums[j]] = true
		}
		i++
		j++
	}
	return false
}

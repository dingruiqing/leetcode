package problem

func getId(x int, w int) int {
	if x < 0 {
		return (x+1)/w - 1
	}
	return x / w
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if t < 0 {
		return false
	}
	d := make(map[int]int)
	w := t + 1
	for i, v := range nums {
		m := getId(v, w)
		if _, ok := d[m]; ok {
			return true
		}
		if j, ok := d[m-1]; ok && v-j < w {
			return true
		}
		if j, ok := d[m+1]; ok && j-v < w {
			return true
		}
		d[m] = v
		if i >= k {
			delete(d, getId(nums[i-k], w))
		}
	}
	return false
}

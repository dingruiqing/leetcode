package problem

func removeElement(nums []int, val int) int {
	c := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] == val {
			continue
		} else {
			nums[c] = nums[j]
			c ++
		}
	}
	return c
}

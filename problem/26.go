package problem

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	index := 0
	val := nums[0]
	for _, v := range nums[1:] {
		if v > val {
			val = v
			nums[index] = val
			index ++
		}
	}
	return index
}
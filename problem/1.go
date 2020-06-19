package problem

func twoSum(nums []int, target int) []int {
	nMap := make(map[int]int)
	for i, v := range nums {
		k, ok := nMap[v]
		if !ok {
			nMap[target - v] = i
		} else {
			return []int{i, k}
		}
	}
	return []int{-1, -1}
}


package problem

func sortColors(nums []int)  {
	n := len(nums)
	zero := -1
	one := 0
	two := n
	for one < two {
		switch nums[one] {
		case 1:
			one ++
		case 2:
			two --
			nums[one] = nums[two]
			nums[two] = 2
		case 0:
			zero ++
			nums[one] = 1
			one ++
			nums[zero] = 0
		}
	}
}
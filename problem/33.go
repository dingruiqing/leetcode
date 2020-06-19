package problem

func search(nums []int, target int) int {
	font := 0
	back := len(nums) - 1
	for font < back {
		mid := (font + back) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[0] < nums[mid] {
			if nums[0] <= target && target < nums[mid] {
				back = mid - 1
			} else {
				font = mid + 1
			}
		} else {
			if nums[mid] <= target && target < nums[len(nums)-1] {
				font = mid + 1
			} else {
				back = mid - 1
			}
		}
	}
	return -1
}

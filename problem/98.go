package problem

func IsValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	nums := []int{}
	travel(root, &nums)
	a := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > a {
			a = nums[i]
		} else {
			return false
		}
	}
	return true
}

func travel(node *TreeNode, nums *[]int) {
	if node == nil {
		return
	}
	travel(node.Left, nums)
	*nums = append(*nums, node.Val)
	travel(node.Right, nums)
}

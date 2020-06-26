package problem

func sortedArrayToBST(nums []int) *TreeNode {
	n := len(nums)
	switch n {
	case 0:
		return nil
	case 1:
		return &TreeNode{Val: nums[0]}
	case 2:
		return &TreeNode{Val: nums[1], Left: sortedArrayToBST(nums[:1])}
	default:
		mid := n >> 1
		return &TreeNode{Val: nums[mid], Left: sortedArrayToBST(nums[:mid]), Right: sortedArrayToBST(nums[mid+1:])}
	}
}

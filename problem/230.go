package problem

func kthSmallest(root *TreeNode, k int) int {
	nums := []int{}
	kth(root, &nums, k)
	return nums[k-1]
}

func kth(root *TreeNode, nums *[]int, k int) {
	if root == nil {
		return
	}
	kth(root.Left, nums, k)
	*nums = append(*nums, root.Val)
	kth(root.Right, nums, k)
}

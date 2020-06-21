package problem

func levelOrderBottom(root *TreeNode) [][]int {
	var nums [][]int
	if root == nil {
		return nums
	}
	find(root, &nums, 0)
	return nums
}

func find(node *TreeNode, nums *[][]int, level int) {
	if len(*nums) < level {
		*nums = append([][]int{}, *nums...)
	}
	(*nums)[len(*nums)-level] = append((*nums)[len(*nums)-level], node.Val)
	level++
	if node.Left != nil {
		find(node.Left, nums, level)
	}
	if node.Right != nil {
		find(node.Left, nums, level)
	}
}

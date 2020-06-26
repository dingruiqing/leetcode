package problem

func LevelOrder(root *TreeNode) [][]int {
	var nums [][]int
	if root == nil {
		return nums
	}
	queue := []*TreeNode{root}
	for i := 0; len(queue) > 0; i++ {
		nums = append(nums, []int{})
		p := []*TreeNode{}
		for j := 0; j < len(queue); j++ {
			node := queue[j]
			nums[i] = append(nums[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}

		}
		queue = p
	}
	return nums
}

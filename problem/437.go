package problem

var times int

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	getSum(root, sum)
	pathSum(root.Left, sum)
	pathSum(root.Right, sum)
	return times
}

func getSum(node *TreeNode, sum int) {
	if node == nil {
		if sum == 0 {
			times++
		}
		return
	}
	getSum(node.Left, sum-node.Val)
	getSum(node.Right, sum-node.Val)
}

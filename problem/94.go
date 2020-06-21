package problem

func InorderTraversal(root *TreeNode) []int {
	var res []int
	if root != nil {
		res = append(res, InorderTraversal(root.Left)...)
		res = append(res, root.Val)
		res = append(res, InorderTraversal(root.Right)...)
	}
	return res
}

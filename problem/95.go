package problem

func GenerateTrees(n int) []*TreeNode {
	res := make([][]*TreeNode, n+1)
	res[0] = []*TreeNode{nil}
	res[1] = []*TreeNode{{Val: 1}}
	for i := 2; i <= n; i++ {
		res[i] = make([]*TreeNode, 0)
		for j := 0; j < i; j++ {
			for _, v := range res[j] {
				left := DeepCopyTree(v, 0)
				for _, k := range res[i-j-1] {
					t := &TreeNode{Val: j + 1}
					t.Left = left
					t.Right = DeepCopyTree(k, j+1)
					res[i] = append(res[i], t)
				}
			}
		}
	}
	return res[n]
}

func DeepCopyTree(node *TreeNode, offset int) *TreeNode {
	if node == nil {
		return nil
	}
	copyTree := &TreeNode{Val: node.Val + offset}
	copyTree.Left = DeepCopyTree(node.Left, offset)
	copyTree.Right = DeepCopyTree(node.Right, offset)
	return copyTree
}

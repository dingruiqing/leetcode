package problem

func isBalanced(root *TreeNode) bool {
	m := make(map[*TreeNode]int)
	return balanced(root, &m)
}

func balanced(root *TreeNode, m *map[*TreeNode]int) bool {
	if root == nil {
		return true
	}
	leftDepth := getDepth(root.Left, m)
	rightDepth := getDepth(root.Right, m)
	d := leftDepth - rightDepth
	if d > 1 || d < -1 {
		return false
	} else {
		return balanced(root.Right, m) && balanced(root.Left, m)
	}
}

func getDepth(node *TreeNode, m *map[*TreeNode]int) int {
	v, ok := (*m)[node]
	if ok {
		return v
	} else {
		if node == nil {
			(*m)[node] = 0
			return 0
		}
		l := getDepth(node.Left, m)
		r := getDepth(node.Right, m)
		if l > r {
			(*m)[node] = l + 1
			return l
		} else {
			(*m)[node] = r + 1
			return r
		}
	}
}

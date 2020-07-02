package problem

//func rob(root *TreeNode) int {
//	if root == nil {return 0}
//	m := make(map[*TreeNode]int)
//	robMap(root, &m)
//	return m[root]
//}
//
//func robMap(node *TreeNode, m *map[*TreeNode]int) int {
//	a, b, c, d, e, f := 0, 0, 0, 0, 0, 0
//	if node.Left != nil {
//		if node.Left.Left != nil {
//			a = getValue(node.Left.Left, m)
//		}
//		if node.Left.Right != nil {
//			b = getValue(node.Left.Right, m)
//		}
//		e = getValue(node.Left, m)
//	}
//	if node.Right != nil {
//		if node.Right.Left != nil {
//			c = getValue(node.Right.Left, m)
//		}
//		if node.Right.Right != nil {
//			d = getValue(node.Right.Right, m)
//		}
//		f = getValue(node.Right, m)
//	}
//	(*m)[node] = max(a + b + c + d + node.Val, f + e)
//	return (*m)[node]
//}
//
//
//func getValue(node *TreeNode, m *map[*TreeNode]int) int {
//	v, ok := (*m)[node]
//	a := 0
//	if ok {
//		a = v
//	} else {
//		a = robMap(node, m)
//	}
//	return a
//}

func rob(root *TreeNode) int {
	res := robInternal(root)
	return max(res[0], res[1])
}

func robInternal(node *TreeNode) []int {
	res := []int{0, 0}
	if node == nil {
		return res
	}
	left := robInternal(node.Left)
	right := robInternal(node.Right)
	res[0] = max(left[0], left[1]) + max(right[0], right[1])
	res[1] = left[0] + right[0] + node.Val
	return res
}

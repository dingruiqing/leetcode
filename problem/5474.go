package problem

func addMap(node *TreeNode, distance int) (map[int]int, int) {
	m := make(map[int]int)
	if node.Left == nil && node.Right == nil {
		m[0] = 1
		return m, 0
	} else if node.Left != nil && node.Right == nil {
		m1, r1 := addMap(node.Left, distance)
		for k, v := range m1 {
			m[k+1] = v
		}
		return m, r1
	} else if node.Left == nil && node.Right != nil {
		m1, r1 := addMap(node.Right, distance)
		for k, v := range m1 {
			m[k+1] = v
		}
		return m, r1
	}
	m1, r1 := addMap(node.Left, distance)
	m2, r2 := addMap(node.Right, distance)
	res := r1 + r2
	for k, v1 := range m1 {
		if k+1 <= distance {
			if v2, ok := m2[k]; ok {
				m[k+1] = v1 + v2
			} else {
				m[k+1] = v1
			}
		}
	}
	for k, v := range m2 {
		if k+1 <= distance {
			if _, ok := m1[k]; !ok {
				m[k+1] = v
			}
		}
	}
	for k1, v1 := range m1 {
		for k2, v2 := range m2 {
			if k1+k2+2 <= distance {
				res += v1 * v2
			}
		}
	}
	return m, res
}

func CountPairs(root *TreeNode, distance int) int {
	if root.Right == nil && root.Left == nil {
		return 0
	} else {
		_, res := addMap(root, distance)
		return res
	}
}

package problem

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	flag := true
	for len(q) != 0 {
		tmp := make([]*TreeNode, 0)
		t := make([]int, 0)
		if flag {
			for i := 0; i < len(q); i++ {
				t = append(t, q[i].Val)
			}
			flag = false
		} else {
			for i := len(q) - 1; i >= 0; i-- {
				t = append(t, q[i].Val)
			}
			flag = true
		}
		for i := 0; i < len(q); i++ {
			if q[i].Left != nil {
				tmp = append(tmp, q[i].Left)
			}
			if q[i].Right != nil {
				tmp = append(tmp, q[i].Right)
			}
		}
		res = append(res, t)
		q = tmp
	}
	return res
}

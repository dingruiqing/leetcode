package problem

import "strconv"

func BinaryTreePaths(root *TreeNode) []string {
	q := []*TreeNode{root}
	str := []string{strconv.Itoa(root.Val)}
	res := []string{}
	for len(q) != 0 {
		tmp := []*TreeNode{}
		tmpStr := []string{}
		for i, v := range q {
			if v.Left != nil {
				tmp = append(tmp, v.Left)
				tmpStr = append(tmpStr, str[i]+"=>"+strconv.Itoa(v.Left.Val))
			}
			if v.Right != nil {
				tmp = append(tmp, v.Right)
				tmpStr = append(tmpStr, str[i]+"=>"+strconv.Itoa(v.Right.Val))
			}
			if v.Right == nil && v.Left == nil {
				res = append(res, str[i])
			}
		}
		q = tmp
		str = tmpStr
	}
	return append(res, str...)
}

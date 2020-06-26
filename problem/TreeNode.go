package problem

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTree(nums []int) *TreeNode {
	n := len(nums)
	m := n
	if n == 0 {
		return nil
	}
	root := &TreeNode{Val: nums[0]}
	if n == 1 {
		return root
	}
	d := -1
	for m > 0 {
		m = (m - 1) / 2
		d++
	}
	tmp := []*TreeNode{}
	sum := 0
	for ; d > 0; d-- {
		v := 2 << (d - 1)
		sum += v
		a := []*TreeNode{}
		for i := 0; i < v; i++ {
			if len(tmp) == 0 {
				a = append(a, &TreeNode{Val: nums[n-sum+i]})
			} else {
				a = append(a, &TreeNode{Val: nums[n-sum+i], Left: tmp[2*i], Right: tmp[2*i+1]})
			}
		}
		tmp = a
	}
	root.Left = tmp[0]
	root.Right = tmp[1]
	return root
}

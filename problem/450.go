package problem

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key == root.Val {
		return deleteTreeNode(root)
	} else if key < root.Val {
		return deleteNode(root.Left, key)
	} else {
		return deleteNode(root.Right, key)
	}
}

func deleteTreeNode(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	for node.Left != nil {
		a := node.Left.Right
		b := node.Left.Left
		node.Left.Right = node
		node.Left.Left = a
		node = b
	}
	return node
}

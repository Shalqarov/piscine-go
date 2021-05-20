package piscine

func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	} else if root.Left == nil {
		return root
	}
	return BTreeMax(root.Left)
}

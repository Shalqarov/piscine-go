package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}
	if elem < root.Data {
		root.Left = BTreeSearchItem(root.Left, elem)
	} else if elem > root.Data {
		root.Right = BTreeSearchItem(root.Right, elem)
	}
	return root
}

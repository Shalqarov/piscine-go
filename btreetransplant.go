package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	branch := BTreeSearchItem(root, node.Data).Parent
	if branch == nil {
		return rplc
	} else if node.Data < root.Data {
		root.Left = rplc
	} else {
		root.Right = rplc
	}
	rplc.Parent = branch
	return root
}

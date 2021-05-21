package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	branch := BTreeSearchItem(root, node.Data)
	parentroot := branch.Parent
	rightroot := branch.Right
	leftroot := branch.Left
	maxLeft := leftroot
	for maxLeft.Right != nil {
		maxLeft = maxLeft.Right
	}
	if maxLeft.Left != nil {
		maxLeft.Parent.Right = maxLeft.Left
	}
	maxLeft.Parent = parentroot
	if maxLeft != leftroot {
		maxLeft.Left = leftroot
	}
	maxLeft.Right = rightroot
	if parentroot == nil {
		return maxLeft
	} else if parentroot.Left != nil {
		parentroot.Left = maxLeft
	} else {
		parentroot.Right = maxLeft
	}
	return root
}

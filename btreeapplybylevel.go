package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	size := BTreeLevelCount(root)
	for i := 1; i <= size; i++ {
		BtreeCurrentLevel(root, i, f)
	}
}

func BtreeCurrentLevel(root *TreeNode, size int, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	if size == 1 {
		f(root.Data)
	} else if size > 1 {
		BtreeCurrentLevel(root.Left, size-1, f)
		BtreeCurrentLevel(root.Right, size-1, f)
	}
}

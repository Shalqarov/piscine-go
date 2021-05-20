package piscine

func BTreeLevelCount(root *TreeNode) int {
	cnt := 0
	if root == nil {
		return cnt
	}
	cnt++
	counterright := cnt + BTreeLevelCount(root.Right)
	counterleft := cnt + BTreeLevelCount(root.Left)
	if counterright > counterleft {
		return counterright
	} else {
		return counterleft
	}
}

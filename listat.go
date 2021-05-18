package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	t := l
	for i := 0; i < pos; i++ {
		if t == nil {
			return nil
		} else {
			t = t.Next
		}
	}
	return t
}

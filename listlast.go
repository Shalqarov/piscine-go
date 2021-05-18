package piscine

func ListLast(l *List) interface{} {
	t := l.Head
	if t == nil {
		return nil
	}
	for t.Next != nil {
		t = t.Next
	}
	return t.Data
}

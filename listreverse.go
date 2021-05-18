package piscine

func ListReverse(l *List) {
	if l.Head == nil {
		return
	}
	current := l.Head
	l.Tail = &NodeL{Data: l.Head.Data, Next: nil}
	prev := l.Tail
	prev = nil
	for current != nil {
		l.Head.Next = current.Next
		current.Next = prev
		prev = current
		current = current.Next
	}
	l.Head = prev
}

package piscine

func ListReverse(l *List) {
	if l.Head == nil {
		return
	}
	prev := &NodeL{Data: l.Head.Data, Next: nil}
	for i := 0; i < ListSize(l); i++ {
		current := l.Head
		for current.Next != nil && current.Next != prev {
			current.Data, current.Next.Data = current.Next.Data, current.Data
			current = current.Next
		}
		prev = current
	}
}

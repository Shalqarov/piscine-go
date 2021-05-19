package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head == nil {
		return
	}
	if l.Head.Data == data_ref {
		for l.Head != nil && l.Head.Data == data_ref {
			l.Head = l.Head.Next
		}
		if l.Head == nil {
			return
		}
	}
	next := l.Head.Next
	prev := l.Head
	for next != nil {
		if next.Data == data_ref {
			for next != nil && next.Data == data_ref {
				next = next.Next
			}
			prev.Next = next
			if next == nil {
				return
			}
		}
		prev = prev.Next
		next = next.Next
	}
}

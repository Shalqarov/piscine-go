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
	current := l.Head.Next
	prev := l.Head
	for current != nil {
		if current.Data == data_ref {
			for current != nil && current.Data == data_ref {
				current = current.Next
			}
			prev.Next = current
			if current == nil {
				return
			}
		}
		current = current.Next
	}
}

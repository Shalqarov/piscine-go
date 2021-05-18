package piscine

func ListPushFront(l *List, data interface{}) {
	if l.Head == nil {
		l.Head = &NodeL{Data: data}
	} else {
		l.Head = &NodeL{Data: data, Next: l.Head}
	}
}

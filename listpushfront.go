package piscine

func ListPushFront(l *List, data interface{}) {
	if l.Head == nil {
		l.Head = &NodeL{Data: data}
	} else {
		temp := l.Head
		temp = temp.Next
		temp.Next = &NodeL{Data: data}
	}
}

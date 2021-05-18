package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	if l.Head == nil {
		l.Head = &NodeL{Data: data}
		l.Head.Next = nil
	} else {
		tail := l.Head
		if tail.Next != nil {
			tail = tail.Next
		}
		tail.Next = &NodeL{Data: data, Next: nil}
	}
}

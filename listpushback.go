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
	} else {
		temp := l.Head
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = &NodeL{Data: data}
	}
}

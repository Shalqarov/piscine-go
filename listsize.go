package piscine

func ListSize(l *List) int {
	if l.Head == nil {
		return 0
	} else {
		cnt := 0
		temp := l.Head
		for temp != nil {
			temp = temp.Next
			cnt++
		}
		return cnt
	}
}

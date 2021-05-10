package piscine

func AppendRange(min, max int) []int {
	var size int = max - min
	if size < 1 {
		return []int(nil)
	}
	mas := []int(nil)
	for i := min; i < max; i++ {
		mas = append(mas, i)
	}
	return mas
}

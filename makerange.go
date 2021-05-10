package piscine

func MakeRange(min, max int) []int {
	var size int = max - min
	if size < 1 {
		return []int(nil)
	}
	mas := make([]int, size)
	for i := min; i < max; i++ {
		mas[i-min] = i
	}
	return mas
}

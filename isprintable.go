package piscine

func IsPrintable(s string) bool {
	temp := []rune(s)
	for i := 0; i < len(temp); i++ {
		if temp[i] <= 31 {
			return false
		}
	}
	return true
}

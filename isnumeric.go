package piscine

func IsNumeric(s string) bool {
	temp := []rune(s)
	for i := 0; i < len(temp); i++ {
		if temp[i] < 48 || temp[i] > 57 {
			return false
		}
	}
	return true
}

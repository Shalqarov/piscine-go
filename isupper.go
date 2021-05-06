package piscine

func IsUpper(s string) bool {
	temp := []rune(s)
	for i := 0; i < len(temp); i++ {
		if temp[i] < 65 || temp[i] > 90 {
			return false
		}
	}
	return true
}

package piscine

func IsLower(s string) bool {
	temp := []rune(s)
	for i := 0; i < len(temp); i++ {
		if temp[i] < 97 || temp[i] > 122 {
			return false
		}
	}
	return true
}

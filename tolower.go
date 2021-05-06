package piscine

func ToLower(s string) string {
	temp := []rune(s)
	for i := 0; i < len(temp); i++ {
		if temp[i] >= 65 && temp[i] <= 90 {
			temp[i] = temp[i] + 32
		}
	}
	str := string(temp)
	return str
}

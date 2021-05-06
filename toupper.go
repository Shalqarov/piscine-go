package piscine

func ToUpper(s string) string {
	temp := []rune(s)
	for i := 0; i < len(temp); i++ {
		if temp[i] >= 97 && temp[i] <= 122 {
			temp[i] = temp[i] - 32
		}
	}
	str := string(temp)
	return str
}

package piscine

func Rot14(s string) string {
	str := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= 65 && s[i] <= 76 || s[i] >= 97 && s[i] <= 108 {
			str += string(rune(s[i] + 14))
		} else if s[i] >= 77 && s[i] <= 90 || s[i] >= 109 && s[i] <= 122 {
			str += string(rune(s[i] - 12))
		} else {
			str += string(s[i])
		}
	}
	return str
}

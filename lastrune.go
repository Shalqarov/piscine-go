package piscine

func LastRune(s string) rune {
	result := []rune(s)
	return result[len(s)-1]
}

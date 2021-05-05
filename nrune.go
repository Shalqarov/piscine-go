package piscine

func NRune(s string, n int) rune {
	result := []rune(s)
	if n < 0 || n >= len(result) {
		return 0
	}
	return result[n]
}

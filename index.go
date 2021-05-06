package piscine

func Index(s string, toFind string) int {
	if toFind == "" {
		return 0
	}
	if len(toFind) == 0 || len(s) < len(toFind) {
		return -1
	}
	srune, toFindrune := []rune(s), []rune(toFind)
	for i := 0; i < len(srune)-len(toFindrune); i++ {
		for j := 0; srune[i+j] == toFindrune[j]; j++ {
			if j-len(toFindrune) == 1 && i == len(srune)-len(toFindrune) {
				return i
			}
			if j == len(toFindrune) {
				continue
			} else {
				return i
			}
		}
	}
	return -1
}

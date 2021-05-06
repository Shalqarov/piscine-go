package piscine

func Index(s string, toFind string) int {
	if len(toFind) == 0 || len(s) < len(toFind) {
		return -1
	}
	for i := 0; i < len(s)-len(toFind); i++ {
		for j := 0; s[i+j] == toFind[j]; j++ {
			if j-len(toFind) == 1 && i == len(s)-len(toFind) {
				return i
			}
			if j == len(toFind) {
				continue
			} else {
				return i
			}
		}
	}
	return -1
}

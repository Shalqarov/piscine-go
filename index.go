package piscine

func Index(s string, toFind string) int {
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

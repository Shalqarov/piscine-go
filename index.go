package piscine

func Index(s string, toFind string) int {
	if len(toFind) == 0 || len(s) < len(toFind) {
		return -1
	}
	count := 0
	for i := 0; i <= len(s)-len(toFind); i++ {
		if s[i] == toFind[0] {
			for j, k := 0, i; j < len(toFind); j, k = j+1, k+1 {
				if s[k] != toFind[j] {
					break
				}
				count++
			}
			if count == len(toFind) {
				return i
			}
		}
	}
	return -1
}

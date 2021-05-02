package piscine

func StrLen(s string) int {
	count := 0
	for index := range s {
		count = index
	}
	return count + 1
}

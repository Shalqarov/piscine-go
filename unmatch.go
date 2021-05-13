package piscine

func Unmatch(a []int) int {
	for i := 0; i < len(a); i++ {
		searching := false
		for j := 0; j < len(a); j++ {
			if a[j] == a[i] && i != j {
				searching = true
			}
		}
		if searching == false {
			return a[i]
		}
	}
	return -1
}

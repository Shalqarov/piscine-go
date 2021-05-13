package piscine

func Unmatch(a []int) int {
	for i := 0; i < len(a); i++ {
		searching := false
		cnt := 0
		for j := 0; j < len(a); j++ {
			if a[j] == a[i] && i != j {
				searching = true
				cnt++
			}
			if cnt == 2 {
				return a[i]
			}
		}
		if searching == false {
			return a[i]
		}
	}
	return -1
}

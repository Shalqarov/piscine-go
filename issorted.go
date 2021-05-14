package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-1; j++ {
			if f(a[j], a[j+1]) > 0 {
				return false
			}
		}
	}
	return true
}

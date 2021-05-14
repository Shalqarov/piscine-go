package piscine

func Any(f func(string) bool, a []string) bool {
	for i := 0; i < len(a); i++ {
		if f(a[i]) == true {
			return true
		}
	}
	return false
}

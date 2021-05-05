package piscine

func Sqrt(nb int) int {
	for i := 1; i < nb/2; i++ {
		if i*i == nb {
			return i
		} else if i*i > nb {
			return 0
		}
	}
	return nb
}

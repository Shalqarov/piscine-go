package piscine

func IterativePower(nb int, power int) int {
	if power == 1 {
		return nb
	}
	if power == 0 {
		return 1
	}
	result := 0
	if power > 0 {
		for i := 0; i <= power; i++ {
			result += nb * nb
		}
	} else {
		for i := 0; i <= power; i++ {
			result += 1 / (nb * nb)
		}
	}
	return result
}

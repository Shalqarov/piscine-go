package piscine

func IterativePower(nb int, power int) int {
	if power == 1 {
		return nb
	}
	if power == 0 {
		return 1
	}
	result := nb
	if power > 0 {
		for i := 0; i < power-1; i++ {
			result *= nb
		}
	} else {
		for i := 0; i < power-1; i++ {
			result *= 1 / (result * nb)
		}
	}
	return result
}

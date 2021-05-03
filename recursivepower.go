package piscine

func RecursivePower(nb int, power int) int {
	if power == 1 {
		return nb
	}
	if power == 0 {
		return 1
	}
	if power < 0 {
		return 0
	} else {
		return nb * RecursivePower(nb, power-1)
	}
}

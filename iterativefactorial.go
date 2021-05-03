package piscine

func IterativeFactorial(nb int) int {
	fuc := 1
	for i := 1; i <= nb; i++ {
		fuc *= i
	}
	return fuc
}

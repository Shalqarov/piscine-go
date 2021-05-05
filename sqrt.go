package piscine

func Sqrt(nb int) int {
	temp := 0
	num := nb
	sqrt := num / 2
	for sqrt != temp {
		temp = sqrt
		sqrt = (num/temp + temp) / 2
		if sqrt == 1 || sqrt*sqrt < nb {
			return 0
		}
	}
	if sqrt*sqrt != nb {
		return 0
	}
	return sqrt
}

package piscine

import "github.com/01-edu/z01"

func PrintNbr(num int) {
	if num < 0 {
		z01.PrintRune('-')
		num *= -1
	} else if num == 0 {
		z01.PrintRune('0')
	}
	var numbers []int32
	for num != 0 {
		d := num % 10
		if d < 0 {
			d *= -1
			num *= -1
		}
		numbers = append(numbers, int32(d))
		num /= 10
	}
	for i := len(numbers) - 1; i >= 0; i-- {
		z01.PrintRune(numbers[i] + 48)
	}
}

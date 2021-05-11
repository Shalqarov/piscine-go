package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func returnInt(num int) int {
	cnt := 0
	if num < 0 {
		z01.PrintRune('-')
		num *= -1
		cnt++
	} else if num == 0 {
		z01.PrintRune('0')
		cnt++
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
		cnt++
	}

	for i := len(numbers) - 1; i >= 0; i-- {
		z01.PrintRune(numbers[i] + 48)
	}
	return cnt
}

func PrintPoints(pointers *point) {
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	returnInt(pointers.x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	returnInt(pointers.y)
}

func main() {
	points := &point{}
	setPoint(points)
	PrintPoints(points)
	z01.PrintRune('\n')
}

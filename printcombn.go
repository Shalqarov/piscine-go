package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		numbers[i] = i
		z01.PrintRune(int32(numbers[i] + 48))
	}
	z01.PrintRune(',')
	z01.PrintRune(' ')
	for {
		numbers[n-1] += 1
		for i := n - 2; i >= 0; i-- {
			if numbers[i+1] >= 10 {
				numbers[i] += 1
			}
		}
		for j := 0; j < n-1; j++ {
			if numbers[j+1] >= 10 {
				numbers[j+1] = numbers[j] + 1
			}
		}
		if numbers[n-1] >= 10 {
			continue
		}
		for i := 0; i < n; i++ {
			z01.PrintRune(int32(numbers[i] + 48))
		}
		if numbers[0] >= 10-n {
			break
		} else {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}

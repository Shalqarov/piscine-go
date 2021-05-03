package piscine

func Fibonacci(index int) int {
	if index < 1 {
		return 0
	} else if index == 1 {
		return 1
	}
	return Fibonacci(index-1) + Fibonacci(index-2)
}

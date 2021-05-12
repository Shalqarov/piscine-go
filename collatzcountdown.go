package piscine

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}
	if start == 1 {
		return 0
	}
	num := start
	for i := 1; ; i++ {
		if num%2 != 0 {
			num = num*3 + 1
		} else {
			num = num / 2
		}
		if num == 1 {
			return i
		}
	}
}

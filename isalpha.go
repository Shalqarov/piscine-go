package piscine

func IsAlpha(s string) bool {
	temp := []rune(s)
	for i := 0; i < len(temp); i++ {
		if temp[i] < 97 || temp[i] > 122 {
			if temp[i] < 65 || temp[i] > 90 {
				if temp[i] < 48 || temp[i] > 57 {
					return false
				}
			}
		}

	}
	return true
}

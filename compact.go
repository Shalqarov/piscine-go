package piscine

func Compact(ptr *[]string) int {
	cnt := 0
	var temp []string
	for i := 0; i < len(*ptr); i++ {
		if (*ptr)[i] != "" {
			temp = append(temp, (*ptr)[i])
			cnt++
		}
	}
	*ptr = temp
	return cnt
}

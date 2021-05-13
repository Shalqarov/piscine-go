package piscine

func Join(strs []string, sep string) string {
	str := strs[0]
	for i := 1; i < len(strs); i++ {
		str += sep + strs[i]
	}
	return str
}

package piscine

func AlphaCount(str string) int {
	var count int = 0
	for _, chars := range str {
		if (chars >= 'A' && chars <= 'Z') || (chars >= 'a' && chars <= 'z') {
			count++
		}
	}
	return count
}

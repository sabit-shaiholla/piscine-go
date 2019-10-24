package piscine

func NRune(s string, n int) rune {
	for index, runes := range s {
		//the index starts from zero, hence +1
		if n == index+1 {
			return runes
		}
	}
	return 0
}

package piscine

func StrRev(s string) string {
	runes := []rune(s)

	var len int
	for a := range runes {
		len = a + 1
	}

	for i, j := 0, len-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

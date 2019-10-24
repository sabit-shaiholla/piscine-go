package piscine

func ToUpper(s string) string {
	array := []rune(s)
	for index, char := range array {
		if char >= 'a' && char <= 'z' {
			// between 'a' and 'A' there are 32 symbols
			array[index] = array[index] - 32
		}
	}
	upString := string(array)
	return upString
}

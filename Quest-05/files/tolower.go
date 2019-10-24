package piscine

func ToLower(s string) string {
	array := []rune(s)
	for index, char := range array {
		if char >= 'A' && char <= 'Z' {
			// between 'a' and 'A' there are 32 symbols
			array[index] = array[index] + 32
		}
	}
	downString := string(array)
	return downString
}

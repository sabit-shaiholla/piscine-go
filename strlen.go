package piscine

func StrLen(str string) int {
	array := []rune(str)
	var len int
	for i := range array {
		len = i + 1
	}
	return len
}

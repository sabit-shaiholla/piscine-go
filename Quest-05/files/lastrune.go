package piscine

func LastRune(s string) rune {
	//convert the string into array of rune
	array := []rune(s)
	var array_len int
	//find the lenght of an array
	for index := range array {
		array_len = index + 1
	}
	//return the last rune in an array
	return array[array_len-1]
}

package piscine

func IsLower(str string) bool {
	//checking whether string is containing smth, if not - false
	if str == "" {
		return false
	}
	array := []rune(str)
	for _, char := range array {
		//checking for lower alphabetic values
		if char >= 'a' && char <= 'z' {
			continue
		} else {
			//if it is not lower alphabetic - then false
			return false
		}
	}
	return true
}

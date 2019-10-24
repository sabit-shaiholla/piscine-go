package piscine

func IsUpper(str string) bool {
	//checking whether string is containing smth, if not - false
	if str == "" {
		return false
	}
	array := []rune(str)
	for _, char := range array {
		//checking for upper alphabetic values
		if char >= 'A' && char <= 'Z' {
			continue
		} else {
			//if it is not upper alphabetic - then false
			return false
		}
	}
	return true
}

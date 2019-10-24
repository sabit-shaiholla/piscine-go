package piscine

func IsNumeric(str string) bool {
	//checking whether string is containing smth, if not - false
	if str == "" {
		return false
	}
	array := []rune(str)
	for _, char := range array {
		//checking for numeric values
		if char >= '0' && char <= '9' {
			continue
		} else {
			//if it is not numeric - then false
			return false
		}
	}
	return true
}

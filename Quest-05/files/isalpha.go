package piscine

func IsAlpha(str string) bool {
	//checking whether string is containing smth, if not - false
	if str == "" {
		return false
	}
	array := []rune(str)
	for _, char := range array {
		//checking for alphanumeric values
		if char >= 'a' && char <= 'z' {
			continue
		} else if char >= 'A' && char <= 'Z' {
			continue
		} else if char >= '0' && char <= '9' {
			continue
		} else {
			//if it is not alphanumeric - then false
			return false
		}
	}
	return true
}

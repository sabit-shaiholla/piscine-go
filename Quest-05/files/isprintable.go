package piscine

func IsPrintable(str string) bool {
	//checking whether string is containing smth, if not - false
	if str == "" {
		return false
	}
	array := []rune(str)
	for _, char := range array {
		//checking for printable values
		// 32 - SPACE, 126 - TILDA
		if char >= ' ' && char <= '~' {
			continue
		} else {
			return false
		}
	}
	return true
}

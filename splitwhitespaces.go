package piscine

func SplitWhiteSpaces(str string) []string {
	//calculating len of str array
	len := 0
	for index, element := range str {
		if (element == ' ' || element == '\n' || element == '\t') && (index != 0 && (str[index-1] != ' ' && str[index-1] != '\n' && str[index-1] != '\t')) {
			len++
		}
	}

	//making dynamic str array
	array := make([]string, len+1)
	index := 0
	tempStr := ""
	for _, element := range str {
		if element == ' ' || element == '\n' || element == '\t' {
			if tempStr != "" {
				array[index] = tempStr
				tempStr = ""
				index++
			}
		} else {
			tempStr = tempStr + string(element)
		}
	}
	if tempStr != "" {
		array[index] = tempStr
	}
	return array
}

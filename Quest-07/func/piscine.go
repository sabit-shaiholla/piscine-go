package piscine

func AppendRange(min, max int) []int {
	var array []int
	for i := min; i < max; i++ {
		array = append(array, i)
	}
	return array
}

func MakeRange(min, max int) []int {
	if min >= max {
		//return an array with nil value
		var array []int = nil
		return array
	} else {
		array := make([]int, max-min)
		index := 0
		for i := min; i < max; i++ {
			array[index] = i
			index++
		}
		return array
	}
}

func ConcatParams(args []string) string {
	str := ""
	var len int
	for index := range args {
		len = index + 1
	}
	for index, element := range args {
		str += element
		if index != len-1 {
			str += "\n"
		}
	}
	return str
}

func SplitWhiteSpaces(str string) []string {
	len := 0
	count := 0
	//finding length of str array
	for index, char := range str {
		if (char == '\n' || char == '\t' || char == ' ') && (index != 0 && str[index-1] != '\n' && str[index-1] != '\t' && str[index-1] != ' ') {
			len++
		}
	}

	strC := ""
	answer := make([]string, len+1)
	for _, element := range str {
		if element != ' ' && element != '\n' && element != '\t' {
			strC += string(element)
		} else if (element == ' ' || element == '\n' || element == '\t') && strC != "" {
			answer[count] = strC
			strC, count = "", count+1
		}
	}
	answer[count] = strC
	return answer
}

package piscine

import "github.com/01-edu/z01"

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

func PrintWordsTables(table []string) {
	for _, word := range table {
		for _, i := range word {
			z01.PrintRune(i)
		}
		z01.PrintRune(10)
	}

}

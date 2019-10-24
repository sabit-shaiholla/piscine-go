package piscine

import (
	"github.com/01-edu/z01"
)

func AlphaCount(str string) int {
	var count int = 0
	for _, chars := range str {
		if (chars >= 'A' && chars <= 'Z') || (chars >= 'a' && chars <= 'z') {
			count++
		}
	}
	return count
}

func SortIntegerTable(table []int) {
	//arr_len := len(table)
	var array_len int
	for index := range table {
		array_len = index
	}
	for i := 0; i < array_len; i++ {
		for j := 0; j < array_len-i; j++ {
			if table[j] > table[j+1] {
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}
}

func PrintNbrInOrder(n int) {
	//checking for the sign
	if n < 0 {
		return
	}
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	//for saving the chars will create a temporary array
	var array []int
	//will iterate through decimal points of a number
	for i := n; i > 0; i = i / 10 {
		remainder := i % 10 //remainder
		array = append(array, remainder)
	}
	//will call the SortIntegerTable function to sort the array in ascending order
	SortIntegerTable(array)
	//will iterate through elements in sorted array and print them
	for _, numbs := range array {
		z01.PrintRune(rune(numbs) + 48)
	}
}

func TrimAtoi(s string) int {
	var sum, ngt int
	for _, a := range s {
		count := 0
		if sum == 0 && a == '-' {
			ngt++
		}
		if a >= '0' && a <= '9' {
			for i := '0'; i < a; i++ {
				count++
			}
			sum = 10*sum + count
		}
	}
	if ngt == 1 {
		return -sum
	}
	return sum
}

func FirstRune(s string) rune {
	for _, runes := range s {
		return runes
	}
	return 0
}

func NRune(s string, n int) rune {
	for index, runes := range s {
		//the index starts from zero, hence +1
		if n == index+1 {
			return runes
		}
	}
	return 0
}

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

func Index(s string, toFind string) int {
	input := []rune(s)
	search := []rune(toFind)
	input_i := 0
	search_i := 0

	for index := range search {
		//to implement range
		index = index
		search_i++
	}
	if search_i == 0 {
		return 0
	}
	for index := range input {
		//to implement range
		index = index
		input_i++
	}
	for index, char := range input {
		if char == search[0] && input_i >= search_i+index-1 {
			var i int = 1
			for j := 1; j < search_i; j++ {
				if search[j] == input[index+j] {
					i++
				}
			}
			if i == search_i {
				return index
			}
		}
	}
	return -1
}

func Compare(a, b string) int {
	if a == b {
		return 0
	} else if a > b {
		return 1
	} else {
		return -1
	}
}

func ToUpper(s string) string {
	array := []rune(s)
	for index, char := range array {
		if char >= 'a' && char <= 'z' {
			// between 'a' and 'A' there are 32 symbols
			array[index] = array[index] - 32
		}
	}
	upString := string(array)
	return upString
}

func ToLower(s string) string {
	array := []rune(s)
	for index, char := range array {
		if char >= 'A' && char <= 'Z' {
			// between 'a' and 'A' there are 32 symbols
			array[index] = array[index] + 32
		}
	}
	downString := string(array)
	return downString
}

func Capitalize(s string) string {
	array := []rune(s)
	counter := 0
	for index := range array {
		index = index
		counter++
	}
	if array[0] >= 'a' && array[0] <= 'z' {
		// between 'a' and 'A' there are 32 symbols
		array[0] = array[0] - 32
	}
	for i := 1; i < counter; i++ {
		//decapitalize alphanumeric values
		if array[i] >= 'A' && array[i] <= 'Z' {
			if array[i-1] >= 'A' && array[i-1] <= 'Z' {
				array[i] = array[i] + 32
			}
			if array[i-1] >= 'a' && array[i-1] <= 'z' {
				array[i] = array[i] + 32
			}
			if array[i-1] >= '0' && array[i-1] <= '9' {
				array[i] = array[i] + 32
			}
		}
		if array[i] >= 'a' && array[i] <= 'z' {
			if array[i-1] >= 'A' && array[i-1] <= 'Z' {
				continue
			}
			if array[i-1] >= 'a' && array[i-1] <= 'z' {
				continue
			}
			if array[i-1] >= '0' && array[i-1] <= '9' {
				continue
			} else {
				array[i] = array[i] - 32
			}
		}
	}
	capString := string(array)
	return capString
}

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

func Concat(str1 string, str2 string) string {
	strConcat := str1 + str2
	return strConcat
}

func BasicJoin(strs []string) string {
	//empty string
	strJoin := ""
	for _, element := range strs {
		strJoin = strJoin + element
	}
	return strJoin
}

func Join(strs []string, sep string) string {
	//empty string
	strJoin := ""
	for index, str := range strs {
		//except for the last element always add separator
		if index > 0 {
			strJoin = strJoin + sep + str
		} else {
			strJoin = strJoin + str
		}
	}
	return strJoin
}

// func PrintNbrBase(nbr int, base string) () {
// 	runes := []rune(base)
// 	for index, char := range base{
// 		if char >= '0' && char <= '9'{

// 		}
// 	}
// }

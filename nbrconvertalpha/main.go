package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Atoi(s string) int {
	var noZero, negative int
	for g, i := range s {
		count := 0
		if g == 0 && (i == '-' || i == '+') {
			if i == '-' {
				negative++
			}
			continue
		}
		if i < '0' || i > '9' {
			return 0
		}
		for j := '0'; j < i; j++ {
			count++
		}
		noZero = noZero*10 + count
	}

	if negative == 1 {
		return -noZero
	} else {
		return noZero
	}
}

func main() {
	funcArgs := os.Args
	//finding the length of the Argument Array
	var len int
	for index := range funcArgs {
		len = index + 1
	}

	if len > 1 {
		charCount := 96 //before 'a'
		//checking for --upper flag
		if os.Args[1] == "--upper" {
			charCount = 64 //before 'A'; this is for capitalizing
		}
		for index, element := range funcArgs {
			if index > 0 {
				// if the element is in 'a'-'z' or 'A'-'Z' range
				charNbr := Atoi(element) + charCount
				if (charNbr >= 97 && charNbr <= 122) || (charNbr >= 65 && charNbr <= 90) {
					z01.PrintRune(rune(charNbr))
				} else {
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}

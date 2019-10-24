package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	funcName := os.Args[1:]
	var len int
	for index := range funcName {
		len = index + 1
	}
	//Sorting Params
	for i := 0; i < len; i++ {
		for j := 0; j < (len - 1 - i); j++ {
			if funcName[j] > funcName[j+1] {
				funcName[j], funcName[j+1] = funcName[j+1], funcName[j]
			}
		}
	}
	//Printing Params
	for _, element := range funcName {
		for _, char := range element {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}

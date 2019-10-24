package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	funcName := os.Args[1:]
	var lenStr int
	for index := range funcName {
		lenStr = index + 1
	}
	//Reversing elements in funcName array
	for i, j := 0, lenStr-1; i < j; i, j = i+1, j-1 {
		funcName[i], funcName[j] = funcName[j], funcName[i]
	}
	for _, element := range funcName {
		for _, char := range element {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}

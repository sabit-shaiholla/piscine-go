package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	funcName := os.Args[1:]
	for _, element := range funcName {
		for _, char := range element {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}

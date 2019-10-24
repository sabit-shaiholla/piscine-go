package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	funcName := os.Args[0]
	for _, char := range funcName {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	funcName := os.Args[0]
	for _,char := range funcName{
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
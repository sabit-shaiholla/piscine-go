package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {
	for _, b := range str {
		z01.PrintRune(b)
	}
}

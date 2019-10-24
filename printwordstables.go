package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(table []string) {
	for _, word := range table {
		for _, i := range word {
			z01.PrintRune(i)
		}
		z01.PrintRune('\n')
	}

}

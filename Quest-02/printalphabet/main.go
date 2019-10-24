package main

import "github.com/01-edu/z01"

func main() {
	var var1 = 'a'
	for var1 <= 'z' {
		z01.PrintRune(var1)
		var1++
	}
	z01.PrintRune('\n')
}

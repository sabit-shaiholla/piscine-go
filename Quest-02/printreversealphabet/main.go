package main

import "github.com/01-edu/z01"

func main() {
	var var1 = 'z'
	for var1 >= 'a' {
		z01.PrintRune(var1)
		var1--
	}
	z01.PrintRune(10)
}

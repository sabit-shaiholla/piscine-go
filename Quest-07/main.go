package main

import (
	piscine "./func"
)

func main() {
	// fmt.Println(piscine.AppendRange(5, 10))
	// fmt.Println(piscine.AppendRange(10, 5))

	// fmt.Println(piscine.MakeRange(5, 10))
	// fmt.Println(piscine.MakeRange(10, 5))

	// test := []string{"Hello", "how", "are", "you?"}
	// fmt.Println(piscine.ConcatParams(test))

	// str := "Hello how are you?"
	// fmt.Println(piscine.SplitWhiteSpaces(str))

	str := "Hello how are you?"
	table := piscine.SplitWhiteSpaces(str)
	piscine.PrintWordsTables(table)
}

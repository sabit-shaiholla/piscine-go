package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	//checking for the sign
	if n < 0 {
		return
	}
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	//for saving the chars will create a temporary array
	var array []int
	//will iterate through decimal points of a number
	for i := n; i > 0; i /= 10 {
		remainder := i % 10 //remainder
		array = append(array, remainder)
	}
	//will call the SortIntegerTable function to sort the array in ascending order
	//SortIntegerTable(array)
	//will iterate through elements in sorted array and print them
	for _, numbs := range array {
		z01.PrintRune(rune(numbs) + 48)
	}
}

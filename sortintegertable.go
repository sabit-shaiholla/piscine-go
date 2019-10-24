//package main
package piscine

//import "fmt"

func SortIntegerTable(table []int) {
	arr_len := len(table)
	for i := 0; i < arr_len; i++ {
		for j := 0; j < (arr_len - 1 - i); j++ {
			if table[j] > table[j+1] {
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}
}

// func main() {
// 	s := []int{7, 89, 9, 10, 11, 68}
// 	SortIntegerTable(s)
// 	fmt.Println(s)
// }

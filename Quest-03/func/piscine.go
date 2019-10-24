package piscine

import "github.com/01-edu/z01"

import "fmt"

func PointOne(n *int) {
	*n = 1
} 

func UltimatePointOne(n ***int){
	***n = 1 
}

func DivMod(a int, b int, div *int, mod *int) {
	var c int
	var d int
	c = a/b
	d = a%b
	*div = c
	*mod = d
}

func UltimateDivMod(a *int, b *int) {
	// d := *a % *b
	// c := *a / *b
	// *b = d
	// *a = c
	*a, *b = *a / *b, *a % *b
}

func PrintStr(str string) {
	for _, b := range str{
		z01.PrintRune(b)
	}
}

func StrLen(str string) int {
	array := []rune(str)
	var len int
	for i := range array{
		len = i + 1
	}
	return len
}

func Swap(a *int, b *int) {
	*a, *b = *b, *a
}

func StrRev(s string) string {
	runes := []rune(s)

	var len int
	for a := range runes{
		len = a + 1
	}

	for i, j := 0, len - 1; i < j; i, j = i+1, j-1{
		runes[i], runes[j] = runes[j], runes[i]
	} 
	return string(runes)
}

func BasicAtoi(s string) int {
	var afterZero, count int
	for _, i := range s{
		for j:='0'; j < i; j++{
			count++
		}
		afterZero = afterZero * 10 + count
		count = 0
	}
	return afterZero
}

func BasicAtoi2(s string) int {
	var afterZero, count int
	for _, i := range s{
		if i < '0' || i > '9'{
			return 0
		}
		for j:='0'; j < i; j++{
			count++
		}
		afterZero = afterZero * 10 + count
		count = 0
	}
	return afterZero
}

func Atoi(s string) int{
	var afterZero, count int
	for _, i := range s{
		if i >= '+' && i <= '9'{
			for j:='0'; j < i; j++{
				count++
			}
			afterZero = afterZero * 10 + count
			count = 0
		} else {
			return 0
		}
	}
	return afterZero
}

func SortIntegerTable(table []int) {
	for i:=0; i < len(table); i++{
		
	}
}
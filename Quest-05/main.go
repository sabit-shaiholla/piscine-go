package main

import (
	"fmt"

	piscine "./func"
)

func main() {
	// str := "Hello 78 World!    4455 /"
	// nb := piscine.AlphaCount(str)
	// fmt.Println(nb)

	//piscine.PrintNbrInOrder(7432023924179416648)
	//piscine.PrintNbrInOrder(0)
	//piscine.PrintNbrInOrder(321)

	// s := "12345"
	// s2 := "str123ing45"
	// s3 := "012 345"
	// s4 := "Hello World!"
	// s5 := "sd+x1fa2W3s4"
	// s6 := "sd-x1fa2W3s4"
	// s7 := "sdx1-fa2W3s4"

	// n := piscine.TrimAtoi(s)
	// n2 := piscine.TrimAtoi(s2)
	// n3 := piscine.TrimAtoi(s3)
	// n4 := piscine.TrimAtoi(s4)
	// n5 := piscine.TrimAtoi(s5)
	// n6 := piscine.TrimAtoi(s6)
	// n7 := piscine.TrimAtoi(s7)

	// fmt.Println(n)
	// fmt.Println(n2)
	// fmt.Println(n3)
	// fmt.Println(n4)
	// fmt.Println(n5)
	// fmt.Println(n6)
	// fmt.Println(n7)

	// z01.PrintRune(piscine.FirstRune("Hello!"))
	// z01.PrintRune(piscine.FirstRune("Salut!"))
	// z01.PrintRune(piscine.FirstRune("Ola!"))
	// z01.PrintRune('\n')

	// z01.PrintRune(piscine.NRune("Hello!", 3))
	// z01.PrintRune(piscine.NRune("Salut!", 2))
	// z01.PrintRune(piscine.NRune("Ola!", 4))
	// z01.PrintRune('\n')

	// z01.PrintRune(piscine.LastRune("Hello!"))
	// z01.PrintRune(piscine.LastRune("Salut!"))
	// z01.PrintRune(piscine.LastRune("Ola!"))
	// z01.PrintRune('\n')

	// fmt.Println(piscine.Index("Hello!", "l"))
	// fmt.Println(piscine.Index("Salut!", "alu"))
	// fmt.Println(piscine.Index("Ola!", "hOl"))

	// fmt.Println(piscine.Compare("Hello!", "Hello!"))
	// fmt.Println(piscine.Compare("Salut!", "lut!"))
	// fmt.Println(piscine.Compare("Ola!", "Ol"))

	//fmt.Println(piscine.ToUpper("Hello! How are you?"))

	// fmt.Println(piscine.Capitalize("Hello! How are you? How+are+things+4you?"))

	// fmt.Println(piscine.IsPrintable("Hello"))
	// fmt.Println(piscine.IsPrintable("Hello\n"))

	toConcat := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(piscine.Join(toConcat, ":"))
}

package main

import (
    "fmt"
    piscine "./func"
)

func main() {
	// arg := 4
	// fmt.Println(piscine.IterativeFactorial(arg))
	
	// arg := 0
	// fmt.Println(piscine.RecursiveFactorial(arg))

	// arg1 := 0
	// arg2 := 0
	// fmt.Println(piscine.IterativePower(arg1, arg2))

	// arg1 := 4
	// arg2 := 3
	// fmt.Println(piscine.RecursivePower(arg1, arg2))

	// arg1 := -5
	// fmt.Println(piscine.Fibonacci(arg1))

	// arg1 := 1
	// arg2 := 9
	// fmt.Println(piscine.Sqrt(arg1))
	// fmt.Println(piscine.Sqrt(arg2))

	// fmt.Println(piscine.IsPrime(5))
	// fmt.Println(piscine.IsPrime(4))

	fmt.Println(piscine.FindNextPrime(9))
	fmt.Println(piscine.FindNextPrime(11))
}


package piscine

func IterativeFactorial(nb int) int {
		result := 1
		if nb >= 1 && nb <= 24{
			for i := 1; i <= nb; i++ {
				result = result * i
			}
		} else {
			result = 0
		}
		return result
}


func RecursiveFactorial(nb int) int {
	if nb == 1 || nb == 0{
		return 1
	} else if nb > 1 {
		return nb * RecursiveFactorial(nb-1)
	}
	return 0
}

func IterativePower(nb int, power int) int {
	result := 1
	if power > 0 {
		for power != 0 {
			result = result * nb
			power = power - 1
		}
	} else if power == 0 {
		result = 1
	} else {
		result = 0
	}
	return result
}

func RecursivePower(nb int, power int) int {
	if power == 0 {
		return 1
	}
	if power >= 1 {
		return nb * RecursivePower(nb, power-1)
	} else {
		return 0
	}
}

func Fibonacci(index int) int {
	if index == 0 {
		return 0
	}
	if index == 1 {
		return 1
	}
	if index < 0 {
		return -1
	} 
	return Fibonacci(index-1) + Fibonacci(index-2)
}



func Sqrt(nb int) int {
	if nb/2 == 1 || nb < 0{
		return 0
	}
	if nb == 1{
		return 1
	}
	var i int = 0
	for i <= nb/2{
		if i * i == nb{
			return i
		}
		i++
	}
	return 0
}

func IsPrime(nb int) bool {
	for i := 2; i <= nb/2; i++ {
        if nb%i == 0 {
            return false
        }
	}
	return nb > 1
}

func FindNextPrime(nb int) int {
	if IsPrime(nb) == true {
		return nb
	}
	for i:=nb; i <= 2147483647; i++{
		if IsPrime(i) == true{
			return i
		}
	}
	return 2
}
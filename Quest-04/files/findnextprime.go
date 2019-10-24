package piscine

// func IsPrime(nb int) bool {
// 	for i := 2; i <= nb/2; i++ {
// 		if nb%i == 0 {
// 			return false
// 		}
// 	}
// 	return nb > 1
// }

func FindNextPrime(nb int) int {
	if IsPrime(nb) == true {
		return nb
	}
	for i := nb; i <= 2147483647; i++ {
		if IsPrime(i) == true {
			return i
		}
	}
	return 2
}

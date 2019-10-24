package piscine

func BasicAtoi(s string) int {
	var afterZero, count int
	for _, i := range s {
		for j := '0'; j < i; j++ {
			count++
		}
		afterZero = afterZero*10 + count
		count = 0
	}
	return afterZero
}

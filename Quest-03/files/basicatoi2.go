package piscine

func BasicAtoi2(s string) int {
	var afterZero, count int
	for _, i := range s {
		if i < '0' || i > '9' {
			return 0
		}
		for j := '0'; j < i; j++ {
			count++
		}
		afterZero = afterZero*10 + count
		count = 0
	}
	return afterZero
}

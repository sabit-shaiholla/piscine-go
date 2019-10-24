package piscine

func Atoi(s string) int {
	var noZero, negative int
	for g, i := range s {
		count := 0
		if g == 0 && (i == '-' || i == '+') {
			if i == '-' {
				negative++
			}
			continue
		}
		if i < '0' || i > '9' {
			return 0
		}
		for j := '0'; j < i; j++ {
			count++
		}
		noZero = noZero*10 + count
	}

	if negative == 1 {
		return -noZero
	} else {
		return noZero
	}
}

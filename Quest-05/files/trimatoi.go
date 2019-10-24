package piscine

func TrimAtoi(s string) int {
	var sum, ngt int
	for _, a := range s {
		count := 0
		if sum == 0 && a == '-' {
			ngt++
		}
		if a >= '0' && a <= '9' {
			for i := '0'; i < a; i++ {
				count++
			}
			sum = 10*sum + count
		}
	}
	if ngt == 1 {
		return -sum
	}
	return sum
}

package piscine

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

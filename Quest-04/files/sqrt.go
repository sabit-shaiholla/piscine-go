package piscine

func Sqrt(nb int) int {
	if nb/2 == 1 || nb < 0 {
		return 0
	}
	if nb == 1 {
		return 1
	}
	var i int = 0
	for i <= nb/2 {
		if i*i == nb {
			return i
		}
		i++
	}
	return 0
}

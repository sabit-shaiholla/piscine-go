package piscine

func DivMod(a int, b int, div *int, mod *int) {
	var c int
	var d int
	c = a / b
	d = a % b
	*div = c
	*mod = d
}

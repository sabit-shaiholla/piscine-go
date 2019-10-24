package piscine

func UltimateDivMod(a *int, b *int) {
	d := *a % *b
	c := *a / *b
	*b = d
	*a = c
}

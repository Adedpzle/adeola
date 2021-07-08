package piscine

func UltimateDivMod(a *int, b *int) {
	aValue := *a
	*a = aValue / *b
	*b = aValue % *b
}

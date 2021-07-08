package piscine

func Swap(a *int, b *int) {
	aValue := *a
	bvalue := *b
	*b = aValue
	*a = bvalue
}

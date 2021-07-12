package piscine

func DivMod(a int, b int, div *int, mod *int) {
	fn0(div, a, b, mod)
}

func fn0(div *int, a int, b int, mod *int) {
	*div = a / b
	*mod = a % b
}

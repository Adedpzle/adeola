package piscine

func IterativeFactorial(nb int) int {
	fct := nb
	if nb > 25 || nb < 0 {
		return nb
	}
	for i := 1; i <= nb; i++ {
		fct = fct * i
	}
	return fct
}

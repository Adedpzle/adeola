package piscine

func IterativeFactorial(nb int) int {
	fct := 1
	if nb > 25 || nb < 0 {
		return 0
	}
	for i := 1; i <= nb; i++ {
		fct = fct * i
	}
	return fct
}

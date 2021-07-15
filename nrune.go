package piscine

func NRune(s string, n int) rune {
	if n <= len(s) && n > 0 {
		ss := []rune(s)
		return ss[n-1]
	}
	return 0
}

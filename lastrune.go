package piscine

func LastRune(s string) rune {
	ss := []rune(s)
	return ss[len(s)-1]
}

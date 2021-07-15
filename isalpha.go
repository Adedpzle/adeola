package piscine

func IsAlpha(s string) bool {
	for _, v := range s {
		if v < '0' || (v > '9' &&
			v < 'A') || (v > 'Z' &&
			v < 'a') || v > 'z' {
			return false
		}
	}
	return true
}

package piscine

func IsUpper(s string) bool {
	for _, v := range s {
		if v < 'A' || v > 'Z' {
			return false
		}
	}
	return true
}

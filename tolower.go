package piscine

func ToLower(s string) string {
	ss := []rune(s)
	for i := 0; i < len(ss); i++ {
		if ss[i] >= 'A' && ss[i] <= 'Z' {
			ss[i] += 32
		}
	}
	return string(ss)
}

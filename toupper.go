package piscine

func ToUpper(s string) string {
	ss := []rune(s)
	for i := 0; i < len(ss); i++ {
		if ss[i] >= 'a' && ss[i] <= 'z' {
			ss[i] -= 32
		}
	}
	return string(ss)
}

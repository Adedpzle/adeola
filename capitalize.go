package piscine

func Capitalize(s string) string {
	ar := []rune(s)

	letra := true

	for i := 0; i < len(s); i++ {
		if IsAlpha(string(ar[i])) && letra {
			if ar[i] >= 'a' && ar[i] <= 'z' {
				ar[i] = 'A' - 'a' + ar[i]
			}
			letra = false
		} else if ar[i] >= 'A' && ar[i] <= 'Z' {
			ar[i] = 'a' - 'A' + ar[i]
		} else if !IsAlpha(string(ar[i])) {
			letra = true
		}
	}
	return string(ar)
}

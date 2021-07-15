package piscine

func Join(strs []string, sep string) string {
	i := ""

	for b, v := range strs {
		if b > 0 {
			i += sep + v
			continue
		}
		i += v
	}

	return i
}

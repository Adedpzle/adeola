package piscine

func TrimAtoi(s string) int {
	bstr := ""
	for _, v := range s {
		if (v >= '0' && v <= '9') || v == '-' {
			if len(bstr) == 0 && v == '-' {
				bstr += string(v)
				continue
			}
			if v != '-' {
				bstr += string(v)
			}
		}
	}
	return Atoi(bstr)
}

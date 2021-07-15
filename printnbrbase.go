package piscine

func NumberToBase(nbr int, base string) string {
	bLen := len(base)
	str := ""
	min := false

	if nbr < 0 {
		min = true
		nbr = nbr * -1
	}

	for {
		d := nbr % bLen
		if d < 0 {
			d *= -1
		}
		str += string(base[d])
		nbr /= bLen

		if nbr == 0 {
			if min {
				str += "-"
			}
			break
		}
	}
	return str
}

func isUnique(s string) bool {
	sr := []rune(s)
	for i := 0; i < len(sr)-1; i++ {
		for j := 0; j < len(sr)-1-i; j++ {
			if sr[j] == sr[j+1] {
				return false
			}
			sr[j], sr[j+1] = sr[j+1], sr[j]
		}
	}
	return true
}

func isSigned(s string) bool {
	for _, v := range s {
		if v == '+' || v == '-' {
			return true
		}
	}
	return false
}

func PrintNbrBase(nbr int, base string) {
	if len(base) < 2 || isSigned(base) || !isUnique(base) {
		PrintStr("NV")
		return
	}
	str := NumberToBase(nbr, base)
	s1 := StrRev(str)
	PrintStr(s1)
}

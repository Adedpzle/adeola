package piscine

func SortString(s *[]rune) {
	for i := 0; i < len((*s))-1; i++ {
		for j := 0; j < len((*s))-1-i; j++ {
			if (*s)[j] > (*s)[j+1] {
				(*s)[j], (*s)[j+1] = (*s)[j+1], (*s)[j]
			}
		}
	}
}

func PrintNbrInOrder(n int) {
	nbr := []rune(ConvNbr(n))
	SortString(&nbr)
	PrintStr(string(nbr))
}

func ConvNbr(n int) string {
	donor := "0123456789"
	str := ""
	min := false

	if n < 0 {
		min = true
		n = n * -1
	}

	for {
		d := n % 10
		if d < 0 {
			d *= -1
		}
		str += string(donor[d])
		n /= 10

		if n == 0 {
			if min {
				str += "-"
			}
			break
		}
	}

	return str
}

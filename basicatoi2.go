package piscine

func BasicAtoi2(s string) int {
	byteArr := []byte(s)

	var total int = 0

	for _, val := range byteArr {
		if val < 48 || val > 57 {
			return 0
		}

		num := int(val) - 48

		total = (total * 10) + num
	}

	return total
}

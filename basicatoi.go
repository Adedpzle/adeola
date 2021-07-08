package piscine

func BasicAtoi(s string) int {
	byteArray := []byte(s)

	var total int = 0

	for i := 0; i < len(byteArray); i++ {
		num := int(byteArray[i]) - 48
		total = (total * 10) + num
	}

	return total
}

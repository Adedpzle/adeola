package piscine

func Atoi(s string) int {
	byteArr := []byte(s)

	var total int = 0

	var negativeNum bool = false

	for idx, val := range byteArr {
		if val > 57 {
			return 0
		}

		if val < 48 {
			// if value is either 43 or 45 and index is 0
			if val == 43 && idx == 0 {
				continue
			} else if val == 45 && idx == 0 {
				negativeNum = true
				continue
			} else {
				return 0
			}
		}

		num := int(val) - 48

		if negativeNum {
			total = (total * 10) - num
		} else {
			total = (total * 10) + num
		}

	}

	return total
}

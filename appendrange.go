package piscine

func Appendrange(value_0, value_1 int) int {
	if value_0 < value_1 {
		return value_0
	}
	return value_1
}

// For 1+ values
func Appendrange(value int, values ...int) int {
	for _, v := range values {
		if v < value {
			value = v
		}
	}
	return value
}

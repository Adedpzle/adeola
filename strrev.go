package piscine

func StrRev(s string) string {
	var reversedStr []byte
	var lastIdx int = len(s) - 1

	for i := lastIdx; i >= 0; i-- {
		reversedStr = append(reversedStr, s[i])
	}

	return string(reversedStr)
}

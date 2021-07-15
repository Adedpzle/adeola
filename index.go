package piscine

func Index(s string, toFind string) int {
	bs := []rune(s)
	stf := []rune(toFind)
	index := -1
	for i := 0; i < len(bs); i++ {
		if len(stf)+i <= len(bs) {
			if s[i:i+len(stf)] == toFind {
				index = i
				break
			}
		}
	}
	return index
}

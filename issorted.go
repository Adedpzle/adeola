package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) > 2 {
		flag := f(a[0], a[1])
		for i := 1; i < len(a)-1; i++ {
			if f(a[i], a[i+1]) > 0 && flag < 0 {
				return false
			}
			if f(a[i], a[i+1]) < 0 && flag > 0 {
				return false
			}
		}
	}
	return true
}

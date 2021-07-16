package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Atoi(s string) int {
	if s == "" {
		return 0
	}

	n := 0

	s0 := s

	if s[0] == '-' || s[0] == '+' {
		s = s[1:]

		if len(s) < 1 {
			return 0
		}
	}

	for _, ch := range []byte(s) {
		ch -= '0'

		if ch > 9 {
			return 0
		}

		n = n*10 + int(ch)
	}

	if s0[0] == '-' {
		n = -n
	}

	return n
}

func ToUpper(s rune) rune {
	if s >= 'a' && s <= 'z' {
		s -= 32
	}

	return s
}

func main() {
	if len(os.Args) <= 1 {
		return
	}
	alpR := []rune("abcdefghijklmnopqrstuvwxyz")

	isUpper := false
	start := 1

	if os.Args[1] == "--upper" {
		isUpper = true
		start = 2
	}

	for _, v := range os.Args[start:] {
		if Atoi(v) > 0 && Atoi(v) < 27 {
			if isUpper {
				z01.PrintRune(ToUpper(alpR[Atoi(v)-1]))
			} else {
				z01.PrintRune(alpR[Atoi(v)-1])
			}
		} else {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}

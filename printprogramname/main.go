package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	runes := []rune(arguments[0])
	var result string
	for i := range runes {
		a := runes[i]
		if a == '/' {
			result = string(a)
		} else {
			result += string(a)
		}
	}
	b := []rune(result)
	for i := range b {
		if i != 0 {
			z01.PrintRune(b[i])
		}
	}
	z01.PrintRune('\n')
}

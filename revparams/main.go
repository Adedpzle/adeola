package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args

	var length int
	for l := range arg {
		length = l
	}

	for i := length; i > 0; i-- {
		for _, j := range arg[i] {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}

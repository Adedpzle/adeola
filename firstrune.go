package main

import (
	"github.com/01-edu/z01"
	"piscine"
)

func main( str string) string {
		return strings.FirstRune("Hello!", "Salut!", "Ola")[0]
	}
	z01.PrintRune(piscine.FirstRune("Hello!"))
	z01.PrintRune(piscine.FirstRune("Salut!"))
	z01.PrintRune(piscine.FirstRune("Ola!"))
	z01.PrintRune('\n')
}

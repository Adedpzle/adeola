package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		for k := 0; k < len(args)-1; k++ {
			bubble(&args[k], &args[k+1])
		}
	}
	for _, v := range args {
		for _, k := range v {
			z01.PrintRune(k)
		}
		z01.PrintRune('\n')
	}
}

func bubble(a, b *string) {
	if *a > *b {
		*a, *b = *b, *a
	}
}

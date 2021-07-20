package main

import (
	"github.com/01-edu/z01"
)

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

type point struct {
	x, y int
}

func main() {
	points := &point{0, 0}

	setPoint(points)

	PrintRuneString("x = ")
	PrintRuneInteger(points.x)
	PrintRuneString(", y = ")
	PrintRuneInteger(points.y)
	PrintRuneString("\n")
}

func PrintRuneInteger(n int) {
	var num_slice []int32

	for n != 0 {
		num_slice = append(num_slice, int32(n%10))
		n /= 10
	}

	for i := len(num_slice) - 1; i >= 0; i-- {
		z01.PrintRune('0' + num_slice[i])
	}
}

func PrintRuneString(s string) {
	for _, v := range s {
		z01.PrintRune(rune(v))
	}
}

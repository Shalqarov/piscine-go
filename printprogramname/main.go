package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	for i, c := range arg[0] {
		if i > 1 {
			z01.PrintRune(c)
		}
	}
	z01.PrintRune('\n')
}

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	for i, c := range arg[0] {
		if i != 0 {
			z01.PrintRune(c)
		}
	}
}

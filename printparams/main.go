package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 1 {
		args := os.Args
		for i := 1; i < len(args); i++ {
			for _, c := range args[i] {
				z01.PrintRune(c)
			}
			z01.PrintRune('\n')
		}
	}
}

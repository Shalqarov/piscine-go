package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 1 {
		args := os.Args

		for i := 1; i < len(args); i++ {
			for j := 1; j < len(args)-1; j++ {
				if args[j] > args[j+1] {
					temp := args[j]
					args[j] = args[j+1]
					args[j+1] = temp
				}
			}
		}

		for i := 1; i < len(args); i++ {
			for _, c := range args[i] {
				z01.PrintRune(c)
			}
			z01.PrintRune('\n')
		}
	}
}

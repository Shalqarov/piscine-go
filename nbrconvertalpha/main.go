package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 1 {
		args := os.Args
		lower := "0abcdefghijklmnopqrstuvwxyz"
		upper := "0ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		result := ""
		if args[1] == "--upper" {
			for i := 2; i < len(args); i++ {
				ch := 0
				for _, el := range args[i] {
					ch = ch*10 + int(el) - 48
				}
				if ch >= 1 && ch <= 26 {
					result += string(upper[ch])
				} else {
					result += " "
				}
			}
		} else {
			for i := 1; i < len(args); i++ {
				ch := 0
				for _, el := range args[i] {
					ch = ch*10 + int(el) - 48
				}
				if ch >= 1 && ch <= 26 {
					result += string(lower[ch])
				} else {
					result += " "
				}
			}
		}
		for _, c := range result {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
	}
}

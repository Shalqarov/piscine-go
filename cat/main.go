package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func printLine(ln string) {
	for _, i := range ln {
		z01.PrintRune(i)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		os.Exit(0)
	}
	if len(args) == 0 {
		_, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			printLine("ERROR: " + err.Error() + "\n")
			os.Exit(1)
		}
		printLine("^C")
	} else {
		for i := 0; i < len(args); i++ {
			filedata, err := ioutil.ReadFile(args[i])
			if err != nil {
				printLine("ERROR: " + err.Error() + "\n")
				os.Exit(1)
			}
			printLine(string(filedata))
		}
	}
}

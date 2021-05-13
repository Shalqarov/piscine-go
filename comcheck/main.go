package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		if args[i] == "01" || args[i] == "galaxy" || args[i] == "galaxy 01" {
			fmt.Println("Alert!!!")
			return
		}
	}
}

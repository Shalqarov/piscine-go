package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("File name missing")
	} else if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
	} else {
		file, err := os.Open(os.Args[1])
		if err != nil {
			return
		}
		fi, err := os.Stat(os.Args[1])
		size := fi.Size()
		finfo := make([]byte, size)
		file.Read(finfo)
		fmt.Print(string(finfo))
	}
}

package main

import (
	"fmt"
	"os"
)

func isnum(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || '9' < s[i] {
			return false
		}
	}
	return true
}

func toint(num string) int {
	result := 0
	for _, n := range num {
		result = result*10 + int(n-48)
	}
	return result
}

func main() {
	args := os.Args[1:]
	if len(args) < 3 || args[0] != "-c" || !isnum(args[1]) {
		os.Exit(1)
	}
	cnt := toint(args[1])
	exstatus := false
	for i := 2; i < len(args); i++ {
		file, err := os.Open(args[i])
		if err != nil {
			fmt.Printf(err.Error() + "\n")
			exstatus = true
		} else {
			fStat, _ := file.Stat()
			data := make([]byte, fStat.Size())
			begin := len(data) - cnt
			if begin < 0 {
				begin = 0
			}
			file.Read(data)
			if len(args) > 3 {
				if i != 2 {
					fmt.Printf("\n")
				}
				fmt.Printf("==> %v <==\n", args[i])
			}
			fmt.Printf(string(data[begin:]))
			file.Close()
		}
	}
	if exstatus == true {
		os.Exit(1)
	}
}

package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("No arguments provided")
		return

	}

	for _, arg := range args {
		fmt.Println(arg)
	}
}

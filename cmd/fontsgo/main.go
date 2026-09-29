package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Println("usage: fontsgo install <font>")
		return
	}

	command := args[0]
	font := args[1]

	switch command {
	case "install":
		// empty

	default:
		fmt.Println("Unknown command:", command)
		os.Exit(1) // exit
	}

	fmt.Println("Command:", command)
	fmt.Println("Font:", font)
}
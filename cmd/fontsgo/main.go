package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	command := args[0]
	font := args[1]

	fmt.Println("Command:", command)
	fmt.Println("Font:", font)
}
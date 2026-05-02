package main

import (
	"fmt"
)

var _ = fmt.Print

func main() {
	fmt.Print("$ ")

	input := make([]byte, 1024)
	n, err := fmt.Scan(input)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	command := string(input[:n])
	fmt.Println(command, ": command not found")
}

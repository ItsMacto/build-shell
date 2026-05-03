package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var _ = fmt.Print

func main() {
	fmt.Print("$ ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	command := strings.TrimSpace(input)
	fmt.Printf("%s: command not found\n", command)
}

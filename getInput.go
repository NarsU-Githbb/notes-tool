package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetInput(encoding string) string {

	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Println("Welcome to the notes tool!\n")
		fmt.Print("Select operation (1-4): \n 1. Show notes. \n 2. Add a note. \n 3. Delete a note. \n 4. Exit. \n")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		fmt.Println()

		switch input {
		case "1", "2", "3", "4":
			return input
		default:
			fmt.Println("Invalid input! Please try again.")
		}
	}
}

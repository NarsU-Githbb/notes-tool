package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetInput(encoding string, message string) {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the notes tool!\n")
	fmt.Print("Select operation (1/2): \n 1. Show notes. \n 2. Add a note. \n 3. Delete a note. \n 4. Exit. \n")

	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		fmt.Println()

		switch input {
		case "1":
			encoding = "Show notes."
		case "2":
			encoding = "Add a note."
		case "3":
			encoding = "Delete a note."
		case "4":
			encoding = "Exit."
		default:
			fmt.Print("Invalid input! Please try again.\n")
			continue
		}
		break
	}
}

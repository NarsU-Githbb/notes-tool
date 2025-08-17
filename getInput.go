package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetMenuChoice() string {
	reader := bufio.NewReader(os.Stdin) // for reading the whole line at a time

	for {
		fmt.Println(ColorBlue + "Welcome to the notes tool!" + ColorReset)
		fmt.Println(ColorYellow + "Select operation (1-5):" + ColorReset)
		fmt.Println(ColorGreen + "1." + ColorReset + " Show notes")
		fmt.Println(ColorGreen + "2." + ColorReset + " Add a note")
		fmt.Println(ColorGreen + "3." + ColorReset + " Delete a note")
		fmt.Println(ColorGreen + "4." + ColorReset + " Exit")
		fmt.Println(ColorRed + "5. Delete all notes" + ColorReset)

		input, _ := reader.ReadString('\n') // reads the user input
		input = strings.TrimSpace(input)
		fmt.Println()

		switch input {
		case "1", "2", "3", "4", "5":
			return input // will send input to the choice in main (line 26)
		default:
			PrintError("Invalid input! Please try again.")
		}
	}
}

// GetLine reads a single line of user input for add and delete note in noteslogic
func GetLine(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func WaitForEnter() {
	fmt.Println("\nPress Enter to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func ShowHelp() {
	fmt.Println(ColorGreen + "NotesTool Help" + ColorReset)
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  ./notestool [COLLECTION]")
	fmt.Println()
	fmt.Println("COLLECTION:")
	fmt.Println("  The name of your note collection (file).")
	fmt.Println("  Example: ./notestool mynotes  → will use mynotes.txt")
	fmt.Println()
	fmt.Println("Menu options:")
	fmt.Println("  1 - Show notes")
	fmt.Println("  2 - Add a note")
	fmt.Println("  3 - Delete a note")
	fmt.Println("  4 - Exit")
	fmt.Println("  5 - Delete all notes")
	fmt.Println()
	fmt.Println("You can type './notestool help' to see this message again.")
}

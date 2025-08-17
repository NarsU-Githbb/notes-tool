package main

import (
	"fmt"
	"strconv"
	"strings"
)

type note struct {
	digits int
	text   string
}

func BuildNote(text string, id int) note { //build a new not
	return note{digits: id, text: text}
}

func ShowNotes(notes []note) {
	if len(notes) == 0 { //if there are no notes println
		fmt.Println(ColorYellow + "No notes yet." + ColorReset)
		return
	}
	fmt.Println(ColorGreen + "Notes:" + ColorReset)
	for i, n := range notes { //loops over all notes and reads them
		if i%2 == 0 { // i%2 alternates colors for each note, %03d creates 3 digits e.g. 001, 002
			fmt.Printf(ColorBlue+"%03d - "+ColorReset+"%s\n", n.digits, n.text)
		} else {
			fmt.Printf(ColorGreen+"%03d - "+ColorReset+"%s\n", n.digits, n.text)
		}
	}
}

func DeleteNote(notes []note) []note {
	if len(notes) == 0 {
		PrintError("No notes to delete.")
		return notes
	}

	ShowNotes(notes)

	input := GetLine("Enter the number of the note to delete (or 0 to cancel): ")
	index, err := strconv.Atoi(input)                  //converts the input string to an integer
	if err != nil || index < 0 || index > len(notes) { // if input is invalid (not a number, negative, or out of range) prints an error and returns slice
		PrintError("Invalid number. No notes deleted.")
		return notes
	}

	if index == 0 {
		return notes // cancel
	}

	index--                                           // convert to 0 based because user interface is 1-based
	notes = append(notes[:index], notes[index+1:]...) //will read notes without index

	for i := range notes { // updates the digits after deletion
		notes[i].digits = i + 1
	}

	return notes
}

func DeleteAllNotes(notes []note) []note {
	confirm := strings.ToLower(GetLine(ColorRed + "Are you sure you want to delete ALL notes? (y/n): " + ColorReset)) //tolower will convert all char to lowercase. Also calls GetLine fo user interface
	if confirm == "y" || confirm == "yes" {
		fmt.Println(ColorGreen + "All notes deleted!" + ColorReset)
		WaitForEnter()
		return []note{} //return empty slice without notes
	}
	fmt.Println(ColorYellow + "Operation cancelled." + ColorReset)
	WaitForEnter()
	return notes
}

func AddNote(notes []note) note {
	userInput := GetLine("Enter your note:\n")
	if strings.TrimSpace(userInput) == "" { //trimspace removes leading and tailing characters to ensure the note hac actual text
		PrintError("Note cannot be empty!")
		return AddNote(notes) // Ask again
	}
	newID := len(notes) + 1 //finds how many notes exits and adds availabe ID
	return BuildNote(userInput, newID)

}

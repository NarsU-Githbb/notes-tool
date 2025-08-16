package main

import (
	"fmt"
	"strconv"
)

type note struct {
	digits int
	text   string
}

func BuildNote(text string) note {
	var noteCounter int
	noteCounter++
	return note{digits: noteCounter, text: text}

}

func ShowNote(n note) {

	fmt.Println("Notes:", fmt.Sprintf("%03d", n.digits))
	fmt.Println(n.text)

}

func DeleteNote(notes []note) []note {
	for _, n := range notes {
		ShowNote(n)
	}

	input := GetInput("Enter the number of the note to delete: ")

	index, err := strconv.Atoi(input)
	if err != nil || index < 1 || index > len(notes) {
		fmt.Println("Invalid number. No notes deleted.")
		return notes
	}

	index--
	return append(notes[:index], notes[index+1:]...)
}

func AddNote() note {
	userInput := GetInput("Enter your note:\n")
	return BuildNote(userInput)

}

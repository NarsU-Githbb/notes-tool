package main

import (
	"fmt"
	"os"
)

func ClearScreen() {
	fmt.Print("\033[H\033[2J") //code only for windows, for linux need to use a different method for clear
}

func main() {
	if len(os.Args) != 2 || os.Args[1] == "help" { //os.Args prediefined slice
		ShowHelp() // this function lies in getinput
		return
	}

	collectionName := os.Args[1] + ".txt"
	notesList := LoadNotes(collectionName) //func in the storage

	for {
		ClearScreen()
		choice := GetMenuChoice() // func in the getinput

		switch choice {

		case "1":
			// ShowNotes(notesList) == changed it because wasnt able to print a full notelist
			ShowNotes(notesList) // func in noteslogic
			WaitForEnter()       // func in the getinput

		case "2":
			newNote := AddNote(notesList) //func in the noteslogic
			notesList = append(notesList, newNote)
			SaveNotes(collectionName, notesList) // func in the storage
		case "3":
			notesList = DeleteNote(notesList) //func in the noteslogic
			SaveNotes(collectionName, notesList)
		case "4":
			fmt.Println(ColorGreen + "Goodbye!" + ColorReset)
			return
		case "5":
			notesList = DeleteAllNotes(notesList)
			SaveNotes(collectionName, notesList)
		default:
			fmt.Println(ColorRed + "Invalid choice. Please select 1-4." + ColorReset)
		}
	}
}

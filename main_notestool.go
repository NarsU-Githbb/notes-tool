package main

import "fmt"

func main() {

notesList := // this where the storage would go
	encoding := getInput()
	var result string

	switch encoding {

	case "1":
		ShowNotes(notesList)
	case "2":
		newNote := AddNote()
		notesList = append(notesList, newNote)
		SaveNotes(collectionName, notesList)
	case "3":
		notesList = DeleteNote(notesList)
		SaveNotes(notesList)
	case "4":
		fmt.Println("Goodbye!")
		return
	}

}

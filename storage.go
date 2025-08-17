package main

import (
	"bufio"
	"fmt"
	"os"
)

// this will load notes
func LoadNotes(filename string) []note {
	file, err := os.Open(filename)
	if err != nil {
		// if note doesnt exist return
		return []note{}
	}
	defer file.Close() // ensure that the file properly closes

	var notes []note                  //store of notes
	scanner := bufio.NewScanner(file) //for reading the file line by line
	counter := 1                      //counter for notes
	for scanner.Scan() {
		notes = append(notes, note{digits: counter, text: scanner.Text()}) //digits and text int located in the notlslogic
		counter++
	}

	if err := scanner.Err(); err != nil {
		PrintError("Error reading notes:" + err.Error())
	}

	return notes
}

func SaveNotes(filename string, notes []note) { //this will help to save notes
	file, err := os.Create(filename)
	if err != nil {
		PrintError("Error saving notes:" + err.Error())
		return
	}
	defer file.Close()

	for _, n := range notes { //loop which saves all of the notes line by line
		fmt.Fprintln(file, n.text)
	}
}

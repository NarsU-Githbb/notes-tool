package main

import "fmt"

type note struct {
	first string
}

func BuildNote(first string) note {

	return note{first: first}

}

func ShowNote(n note) {

	fmt.Println("Notes:")
	fmt.Println(n.first)

}

func DeleteNote() {

}

func AddNote() {
	userInput := GetInput("Enter your note:\n")
	n := BuildNote(userInput)
	ShowNote(n)
}

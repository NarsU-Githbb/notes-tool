package main

import (
	"fmt"
)

func PrintError(msg string) { //when you call this command it isn't necessary to type colorRed/reset for error message
	fmt.Println(ColorRed + "Error: " + msg + ColorReset)
}

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
)

package main

import (
	"os"
)

func main() {
	arg := os.Args[1:]

	// Check if user has included the correct number of arguments (outside program name)
	if len(arg) != 1 {
		// Handle instances of no arguments
		if len(arg) == 0 {
			PrintLine("Please include the name of a text file as an argument")
			return
		} else {
			// Handle more than one argument
			PrintLine("You have way too many arguments!\nOnly include one text file as argument")
			return
		}
	}

	// Store contents of arg in 'file'
	file := arg[0]

	// End program in the instance of non-text files
	if !CheckExtension(file) {
		PrintLine("Wrong file format!\nThe file parsed as an argument must be a '.txt' file.")
		return
	}
}

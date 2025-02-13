package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
-help
-add note
-remove note
-persist all current notes to json
-load in json notes from file (expected format, otherwise fails)
//this will be all of the options supported as this is mainly an exercise in learning language
//basics and such
*/
func main() {
	fmt.Println("Hello and welcome to this simple to-do app!")

	reader := bufio.NewReader(os.Stdin)

	notes := make(map[int]string)
	var counter int = 0

	//this will then block and await input, the following options are:
	/*
	   -help
	   -add note
	   -remove note
	   -persist all current notes to json
	   -load in json notes from file (expected format, otherwise fails)
	   //this will be all of the options supported as this is mainly an exercise in learning language
	   //basics and such
	*/
	for {
		fmt.Print(">>>")
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		//compare the command to the built in options
		if command == "help" {

			fmt.Println("Command Available: ")
			fmt.Println("Help - Prints this message")
			fmt.Println("Remove - Prints all current tasks with reference ID's \n, Then prompts for a comma seperated list on which ones to delete")
			fmt.Println("Load - Prompts for a path to a file, expects values to be seperated by '/\n' ")
			fmt.Println("Add - Prompts for a note to be taken")
			fmt.Println("List - Lists all current notes in a formatted fashion")

		} else if command == "Add" {
			//temp to appease compiler?
			fmt.Println("Please enter a note")
			note, err := reader.ReadString('\n')
			if err != nil {
				fmt.Errorf("something happened!!")
				os.Exit(1)
			}

			counter++
			notes[counter] = note
		} else if command == "Remove" {
			fmt.Println("Enter the number of the note you would like to delete")
			printNotes(&notes)
			numberToDelete, err := reader.ReadString('\n')
			numberToDelete = strings.TrimSpace(numberToDelete) // Remove newline characters
			if err != nil {
				fmt.Errorf("something happened!!")
				os.Exit(1)
			}
			//convert string to int
			numberToDeleteIntified, err := strconv.Atoi(numberToDelete)
			fmt.Println(numberToDeleteIntified)
			if err != nil {
				fmt.Println("Invalid number entered")
			}
			delete(notes, numberToDeleteIntified)

		} else if command == "Load" {

		} else if command == "List" {
			printNotes(&notes)
		} else if command == "Clear" {
			for i := 0; i < 10; i++ {
				fmt.Println("")
			}
		} else if command == "Exit" {
			os.Exit(0)
		} else {
			fmt.Println("Invalid entry")
		}

	}

}

func printNotes(notes *map[int]string) {

	for key, value := range *notes {
		fmt.Println("note ", key, ": ", value)
	}

}

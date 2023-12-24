package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/thegera4/dummy_note_app/note"
	"github.com/thegera4/profit_calculator/press_enter_to_exit"
)

func main() {
	title, content := getNoteData()

	newNote, err := note.New(title, content)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	newNote.Display()

	newNoteErr := newNote.SaveToFile()
	if newNoteErr != nil {
		fmt.Println("Saving the note failed!")
		return
	}

	fmt.Println("Note saved!")
	press_enter_to_exit.WaitForEnter()
}

func getNoteData() (string, string) {
	title := getUserInput("Enter a title: ")
	content := getUserInput("Note content: ")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	bufioReader := bufio.NewReader(os.Stdin) //for long text with spaces instead of Scan
	input, err := bufioReader.ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
		return ""
	} 
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")
	return input	
}
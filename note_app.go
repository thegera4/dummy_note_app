package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/thegera4/example/note"
)

func main() {
	title, content := getNoteData()
	newNote, err := note.New(title, content)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	newNote.Display()
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
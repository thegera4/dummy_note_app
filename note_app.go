package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/thegera4/dummy_note_app/note"
	"github.com/thegera4/dummy_note_app/todo"
	"github.com/thegera4/profit_calculator/press_enter_to_exit"
)

//convention in go is that if you only have one method in the interface, the name should be
//the methods name + r at the end
type saveToFiler interface { //a contract that ensures that every struct that implements this interface
	SaveToFile() error //has a SaveToFile() method
}

/*type displayer interface {
	Display()
}*/

//embedded interface
type outputtable interface {
	saveToFiler
	Display()
}

/*Alternative: new interface with both methods
type outputtable interface {
	SaveToFile() error
	Display()
}*/

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Enter a todo: ")

	newTodo, err := todo.New(todoText)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	newNote, err := note.New(title, content)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	err = outputData(newTodo)
	if err != nil {
		return
	}

	outputData(newNote)

	press_enter_to_exit.WaitForEnter()
}

//this function accepts any type of value: strings, ints, floats, etc
func printSomething(value interface{}){ //alternative, instead of interface{} you can use "any"
	intVal, ok := value.(int) //check if the value is an int
	if ok {
		intVal += 10
	}
	floatVal, ok := value.(float64) //check if the value is an float64
	if ok {
		floatVal += 10.5
	}
	stringVal, ok := value.(string) //check if the value is an string
	if ok {
		fmt.Println(stringVal + " is a string!")
		return
	}
	/*switch value.(type) { //alternative with special switch to check the type of the value and do something accordingly
		case string:
			fmt.Println("It's a string!")
		case int:
			fmt.Println("It's an int!")
		case float64:
			fmt.Println("It's a float!")
		default:
			fmt.Println("I don't know what it is!") //you can omit the deafault and nothing will happen
	}*/
	//fmt.Println(value)
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saveToFiler) error {
	err := data.SaveToFile()
	if err != nil {
		fmt.Println("Saving failed!")
		return err
	}
	fmt.Println("Data saved!")
	return nil
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
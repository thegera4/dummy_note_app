package todo

import (
	"errors"
	"fmt"
	"os"
	"encoding/json"
)

type Todo struct {
	Text string `json:"text"`
}

func New(content string) (Todo, error) {
	if content == "" {
		fmt.Println("Please enter a valid input")
		return Todo{}, errors.New("Invalid input")
	}

	return Todo{
		Text: content,
	}, nil
}

func (t Todo) Display() {
	fmt.Println("Your todo: ", t.Text)
}

func (t Todo) SaveToFile() error {
	fileName := "todo.json"
	json, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}
package note

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
	"encoding/json"
)

type Note struct {
	Title     string `json:"title"` //struct tag to change the name of the field in the json file
	Content   string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		fmt.Println("Please enter a valid input")
		return Note{}, errors.New("Invalid input")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (n Note) Display() { //no pointer because we are not editing any values
	fmt.Printf("Your note titled %v has the following content:\n\n%v\n\n", n.Title, n.Content)
}

func (n Note) SaveToFile() error { //no pointer because we are not editing any values
	fileName := strings.ReplaceAll(n.Title, " ", "_") //replace spaces with underscores
	fileName = strings.ToLower(fileName) + ".json" //convert to lowercase and add .json extension
	json, err := json.Marshal(n) //convert data to json
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}
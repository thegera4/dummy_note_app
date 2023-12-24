package note

import (
	"time"
	"errors"
	"fmt"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		fmt.Println("Please enter a valid input")
		return nil, errors.New("Invalid input")
	}

	return &Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil
}

func (n *Note) Display() {
	fmt.Printf("Your note titled %v has the following content:\n\n%v\n\n", n.title, n.content)
}
package openai

import "fmt"

type Error struct {
	StatusCode int    `json:"-"`
	Message    string `json:"message"`
	Type       string `json:"type"`
}

func (e Error) Error() string {
	return fmt.Sprintf("openai: %s (%s)", e.Message, e.Type)
}

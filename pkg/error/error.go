package errors

import (
	"fmt"
)

type Error struct {
	Message string `json:"message"`
}

func NewError(msg string) *Error {
	return &Error{
		Message: msg,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("Message :%s ", e.Message)
}

func (e *Error) EditMessage(message string) {
	e.Message = message
}

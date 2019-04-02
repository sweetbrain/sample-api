package common

import "fmt"

type ErrorMessage struct {
	Code 		int 	`json:"code"`
	Message 	string 	`json:"message"`
}

func NewError(code int, message string) ErrorMessage {
	return ErrorMessage{
		Code: code,
		Message: message,
	}
}

func (e ErrorMessage) Error() string {
	return fmt.Sprintf("code=[%d] message=[%s]", e.Code, e.Message)
}
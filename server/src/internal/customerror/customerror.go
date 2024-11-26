package customerror

import "fmt"

type CustomError struct {
	UserMessage string
	LogMessage  string
}

func New(message string, errs ...error) CustomError {
	if len(errs) > 0 {
		return CustomError{
			UserMessage: message,
			LogMessage:  fmt.Sprintf("%s: %s", message, errs[0].Error()),
		}
	} else {
		return CustomError{
			UserMessage: message,
			LogMessage:  message,
		}
	}
}

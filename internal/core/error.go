package core

import "strings"

type appError struct {
	Type    string `json:"type"`
	Messages  []string `json:"messages"`
}
func (e *appError) Error() string {
	if len(e.Messages) == 0 {
		return ""
	}
	return strings.Join(e.Messages, "; ")
}

func (e *appError) Code() string {
	return e.Type
}

func NewValidationError(errors []string) *appError {
	return &appError{
		Type:   "validation_error",
		Messages: errors,
	}
}

func NewDatabaseConnectionError(errors []string) *appError {
	return &appError{
		Type:   "database_connection_error",
		Messages: errors,
	}
}

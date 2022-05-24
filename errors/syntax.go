package errors

import "fmt"

func Syntax(message string) error {
	return fmt.Errorf("SyntaxError: %s", message)
}

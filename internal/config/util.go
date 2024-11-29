package config

import "fmt"

func formatError(msg string, err error) error {
	return fmt.Errorf("%s - err: %v", msg, err)
}

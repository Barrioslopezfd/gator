package main

import (
	"context"
	"errors"
	"fmt"
)

func handleLogin(s *state, cmd command) error {
    if len(cmd.arguments) < 1 {
	    return errors.New("Not enough arguments")
    }
    if len(cmd.arguments) > 1 {
	    return fmt.Errorf("Too many arguments received. Excepted 1 argument containing username, received:  %d", len(cmd.arguments))
    }
    _, err := s.db.GetUser(context.Background(), cmd.arguments[0])
    if err != nil {
	return fmt.Errorf("User not registered")
    }
    s.conf.SetUser(cmd.arguments[0])
    fmt.Println("User setted as: ", cmd.arguments[0])
    return nil
}


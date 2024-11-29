package main

import (
	"context"
	"fmt"
	"time"
	"github.com/Barrioslopezfd/gator/internal/database"
	"github.com/google/uuid"
)

func handleRegister(s *state, cmd command) error {
    if len(cmd.arguments) < 1 {
	    return fmt.Errorf("Not enough arguments on register")
    }
    if len(cmd.arguments) > 1 {
	    return fmt.Errorf("Too many arguments received. Excepted 1 argument containing username, received:  %d", len(cmd.arguments))
    }

    _, err := s.db.GetUser(context.Background(), cmd.arguments[0])
    if err == nil {
	return fmt.Errorf("User already registered")
    }


    user := database.CreateUserParams {
	ID:	    uuid.New(),
	CreatedAt:  time.Now(),
	UpdatedAt:  time.Now(),
	Name:	    cmd.arguments[0],
    }
    usr, err:= s.db.CreateUser(context.Background(), user)
    if err != nil {
	return err
    }

    s.conf.SetUser(cmd.arguments[0])

    fmt.Println("User created\n", usr)

    return nil
}



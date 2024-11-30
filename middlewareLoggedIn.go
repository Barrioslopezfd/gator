package main

import (
	"context"

	"github.com/Barrioslopezfd/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		currentUser, err := s.db.GetUser(context.Background(), s.conf.CurrentUserName)
		if err != nil {
			return err
		}
		err = handler(s, cmd, currentUser)
		if err != nil {
			return err
		}
	return nil
	}
}

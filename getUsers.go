package main

import (
	"context"
	"fmt"
)

func getUsers(s *state, _ command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	for _, usr := range users {
		if usr.Name == s.conf.CurrentUserName {
			fmt.Println(usr.Name, "(current)")
		} else {
			fmt.Println(usr.Name)
		}
	}
	return nil
}

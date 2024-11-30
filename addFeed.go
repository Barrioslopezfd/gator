package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Barrioslopezfd/gator/internal/database"
	"github.com/google/uuid"
)

func handleAddFeed(s *state, cmd command) error {
	if len(cmd.arguments) < 2 {
		return fmt.Errorf("Not enough arguments, expected 2 but received %d", len(cmd.arguments))
	}
	currentUser, err := s.db.GetUser(context.Background(), s.conf.CurrentUserName)
	if err != nil {
		return err
	}
	newFeed := database.CreateFeedParams {
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.arguments[0],
		Url:       cmd.arguments[1],
		UserID:    currentUser.ID,
	}

	nf, err := s.db.CreateFeed(context.Background(), newFeed)
	if err != nil {
		return err
	}

	fmt.Println(nf.ID)
	fmt.Println(nf.CreatedAt)
	fmt.Println(nf.UpdatedAt)
	fmt.Println(nf.Name)
	fmt.Println(nf.Url)
	fmt.Println(nf.UserID)
	
	return nil
}

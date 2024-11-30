package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Barrioslopezfd/gator/internal/database"
	"github.com/google/uuid"
)


func handleFollow(s *state, cmd command, currentUser database.User) error {
	if len(cmd.arguments) > 1 {
		return fmt.Errorf("Too many arguments. Expected 1, received %d.", len(cmd.arguments))
	}
	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.arguments[0])
	if err != nil {
		return err
	}
	newFeedFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    currentUser.ID,
		FeedID:    feed.ID,
	}
	
	feedFollow, err := s.db.CreateFeedFollow(context.Background(), newFeedFollow)
	if err != nil {
		return err
	}
	for i := range feedFollow {
		fmt.Println(feedFollow[i].FeedName)
		fmt.Println(feedFollow[i].UserName)
	}

	return nil
}

package main

import (
	"context"
	"fmt"

	"github.com/Barrioslopezfd/gator/internal/database"
)

func handleFollowing(s *state, _ command, currentUser database.User) error {
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), currentUser.Name)
	if err != nil {
		return err
	}
	for i := range feeds {
		fmt.Println(feeds[i].Name)
	}
	return nil
}

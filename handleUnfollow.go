package main

import (
	"context"
	"fmt"

	"github.com/Barrioslopezfd/gator/internal/database"
)


func handleUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments)>1 {
		return fmt.Errorf("Too many arguments. Expected 2, received %d", len(cmd.arguments))
	}
	toDelete := database.DeleteFeedFollowByUrlParams{
		Url:	cmd.arguments[0],
		UserID: user.ID,
	}
	err := s.db.DeleteFeedFollowByUrl(context.Background(), toDelete)
	if err != nil {
		return err
	}
	return nil
}

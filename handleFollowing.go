package main

import (
	"context"
	"fmt"
)

func handleFollowing(s *state, _ command) error {
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), s.conf.CurrentUserName)
	if err != nil {
		return err
	}
	for i := range feeds {
		fmt.Println(feeds[i].Name)
	}
	return nil
}

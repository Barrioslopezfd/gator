package main

import (
	"context"
	"fmt"
)

func handleFeeds(s *state, _ command) error {
	feed, err := s.db.GetAlmostFeed(context.Background())
	if err != nil {
		return err
	}
	for i := range feed {
		fmt.Println(feed[i].Name)
		fmt.Println(feed[i].Url)
		fmt.Println(feed[i].Name_2)
	}
	return nil
}

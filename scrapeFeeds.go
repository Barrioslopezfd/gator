package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Barrioslopezfd/gator/internal/database"
)


func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	markedFeed := database.MarkFeedFetchedParams{
		UpdatedAt:	time.Now(),
		ID:			nextFeed.ID,
	}
	s.db.MarkFeedFetched(context.Background(), markedFeed)

	feed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return err
	}
	for i := range feed.Channel.Item {
		fmt.Println(feed.Channel.Item[i].Title)
	}

	return nil
}
